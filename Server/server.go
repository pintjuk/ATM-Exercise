package main

import (
	"../protocol"
	"bytes"
	"encoding/binary"
	"fmt"
	//"io"
	"net"
	"sync"
	"time"
)

const (
	noLangUppdate = "not"
)

type intChan struct {
	id      int32
	channel chan string
}

func construct() (talk func(net.Conn), update func(net.Conn), rprinting <-chan string) {
	// this lame and is not in use curently 
	printing := make(chan string, 5)
	rprinting = printing

	// map of all client update <-> talk chanals and mutex
	clConnTrack := make(map[int32]chan string)
	clCTMtx := new(sync.Mutex)

	// chanal to triger update of all clients 
	updateAll := make(chan bool)

	// loads language set and mutex
	langs := MakeDefoultLangs()
	langMtx := new(sync.Mutex)

	// generates unique IDs for clients and mutex 
	var cc int32 = 0
	ccMtx := new(sync.Mutex)

	// Accounts and mutex
	accounts := makeDumyAccounts()
	accMtx := new(sync.Mutex)

	// simple mass update with updateAll channal
	go func() {
		for _ = range updateAll {
			for stuff := range clConnTrack {
				clConnTrack[stuff] <- noLangUppdate
			}
		}
	}()

	// server UI logic
	go func() {
		for {
			fmt.Println("\t1. Eddit lang\n\t2. update all clients\n\n> ")
			var in uint8
			fmt.Scanln(&in)
			switch in {
			case 1:
				fmt.Println("Exsisting languages: ")
				langMtx.Lock()
				for k, _ := range langs {
					fmt.Println(fmt.Sprintf("\t%s", k))
				}
				langMtx.Unlock()

				fmt.Print("Enter lang code,\nunique creates new lang.\n\n>")
				var in string
				fmt.Scanln(&in)
				_, ok := langs[in]
				if ok {
					fmt.Println("ADD:")
					langMtx.Lock()
					langToEdit := in
					fmt.Println(langs[langToEdit][ADD])
					langMtx.Unlock()
					fmt.Println("\t1. Change addvertisment")
					var in uint8
					fmt.Scanln(&in)
					if in == 1 {
						fmt.Print("Enter new addvertisment text (terminate with \"end\")\n> ")
						var newadd string
						for {
							var l string
							fmt.Scanln(&l)
							if l != "end" {
								newadd = fmt.Sprintf("%s\n%s", newadd, l)
								continue
							}
							break
						}

						langMtx.Lock()
						langs[langToEdit][ADD] = newadd
						langMtx.Unlock()
						updateAll <- true
					}
				}

				continue
			case 2:
				continue
			default:
				fmt.Println("input error")
				continue
			}

		}
	}()

	// Saifly access languages
	mkUpdate := func(lang string) []byte {
		langMtx.Lock()
		defer langMtx.Unlock()
		return protocol.MakeUpdate(langs[lang][ADD], langs[lang])
	}

	// all cliant logic goes here
	talk = func(conn net.Conn) {
		var id func() int32
		func() {
			var idtemp int32
			arr := make([]byte, 4)
			_, err := conn.Read(arr)
			if err != nil {
				fmt.Println(err)
			}
			buf := bytes.NewBuffer(make([]byte, 0, 4))
			buf.Write(arr)
			binary.Read(buf, binary.BigEndian, &idtemp)
			id = func() int32 {
				return idtemp
			}
		}()

		var updateChan chan string

		for {
			clCTMtx.Lock()
			var ok bool
			updateChan, ok = clConnTrack[id()]
			clCTMtx.Unlock()
			if ok {
				break
			} else {
				time.Sleep(time.Millisecond * 100)
			}
		}

		for {
			time.Sleep(time.Millisecond * 10)
			conn.Write(protocol.MakePrintCMD(INIT, nil, protocol.CLEAR_FlAG, protocol.SEND_UINT))
			//for {
			buff := make([]byte, 10)
			n, err := conn.Read(buff)
			if err != nil {
				return
			}
			l, _ := protocol.DecodeUintMSG(buff[:n])
			if l == 1 {
				conn.Write(protocol.MakePrintCMD(ENTACC, nil, protocol.CLEAR_FlAG, protocol.SEND_UINT))
				buff = make([]byte, 10)
				n, err := conn.Read(buff)
				if err != nil {
					return
				}
				accnumber, _ := protocol.DecodeUintMSG(buff[:n])
				accMtx.Lock()
				account, ok := accounts[accnumber]
				accMtx.Unlock()
				if ok {
					for failcount := 1; failcount < 4; failcount++ {
						conn.Write(protocol.MakePrintCMD(ENTCODE, nil, protocol.CLEAR_FlAG, protocol.SEND_UINT))
						buff = make([]byte, 10)
						n, err = conn.Read(buff)
						if err != nil {
							return
						}
						code, _ := protocol.DecodeUintMSG(buff[:n])

						if account.login(uint8(code)) {
							conn.Write(protocol.MakePrintStringCMD(SUCLOGIN, account.name, protocol.CLEAR_FlAG))
							time.Sleep(time.Millisecond * 10)
							conn.Write(protocol.MakePrintFloatCMD(BALANCE, account.balance))
							time.Sleep(time.Second * 20)
							break
						} else {
							conn.Write(protocol.MakePrintUintCMD(ERR_CODE, uint64(failcount), protocol.CLEAR_FlAG))
							time.Sleep(time.Second * 2)
						}
					}

					accMtx.Lock()
					accounts[accnumber] = account
					accMtx.Unlock()

				} else {
					conn.Write(protocol.MakePrintCMD(ERR_ENTACC, nil, protocol.CLEAR_FlAG))
					time.Sleep(time.Second * 2)
				}

			}
			if l == 2 {
				conn.Write(protocol.MakePrintCMD(LANG, nil, protocol.CLEAR_FlAG, protocol.SEND_UINT))
				buff = make([]byte, 10)
				n, err := conn.Read(buff)
				if err != nil {
					return
				}
				l, _ = protocol.DecodeUintMSG(buff[:n])
				switch l {
				case 1:
					updateChan <- "en"
					break
				case 2:
					updateChan <- "sv"
					break
				case 3:
					updateChan <- "ch"
					break
				}
				//fmt.Println(curentLang)
			}
		}
		time.Sleep(time.Minute * 1)
	}

	// update cliant language set
	update = func(conn net.Conn) {
		updatechan := make(chan string)
		ccMtx.Lock()
		cc++
		var id func() int32
		func() {
			id = func() int32 {
				var idtemp int32 = cc
				return idtemp
			}
		}()
		ccMtx.Unlock()

		buf := bytes.NewBuffer(make([]byte, 0, 4))
		binary.Write(buf, binary.BigEndian, id())
		buf.Bytes()
		conn.Write(buf.Bytes())

		clCTMtx.Lock()
		clConnTrack[id()] = updatechan
		clCTMtx.Unlock()

		time.Sleep(time.Millisecond * 10)

		lang := "en"
		go func() { updatechan <- lang }()
		for langtemp := range updatechan {
			//update language if it is not a mass update
			if langtemp != noLangUppdate {
				lang = langtemp
			}
			conn.Write(mkUpdate(lang))
			time.Sleep(time.Second * 3)

		}
	}
	return
}

func main() {
	talk, update, _ := construct()

	go server(talk, ":9000")
	go server(update, ":9001")

	// stuped thing hack 
	// TODO: fix this stupid hack
	w := new(sync.WaitGroup)
	w.Add(1)
	w.Wait()
}

func server(serv func(net.Conn), port string) {
	ln, err := net.Listen("tcp", port)
	wazzup(err)
	for {
		conn, err := ln.Accept()
		wazzup(err)
		go func(conn net.Conn) {
			serv(conn)
			conn.Close()
		}(conn)
	}
}

func wazzup(err error) {
	if err != nil {
		panic(err)
	}
}
