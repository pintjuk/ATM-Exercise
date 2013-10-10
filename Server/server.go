package main

import (
	"../protocol"
	"bytes"
	"fmt"
	"io"
	"net"
	"time"
)

func construct() (talk func(net.Conn), update func(net.Conn), rprinting <-chan string) {
	printing := make(chan string, 5)
	rprinting = printing
	langs := MakeDefoultLangs()
	curentLang := "en"

	mkUpdate := func() []byte {
		return protocol.MakeUpdate(langs.Lang(curentLang)[ADD], langs.Lang(curentLang))
	}

	talk = func(conn net.Conn) {

		printing <- fmt.Sprintf("conection from: %s accepted", conn.RemoteAddr().String())
		var buff bytes.Buffer
		b := make([]byte, 100)
		for {
			n, err := conn.Read(b)
			buff.Write(b[:n])

			if err != io.EOF {
				/*for buff.Len() > 0 {
					r, _, e := buff.ReadRune()
					if e != nil {
						continue
					}
					printing <- fmt.Sprintf(" > %s", r)
				}*/
				//s, _ := buff.ReadString(byte('\'))
				//printing <- fmt.Sprintf(" > %b %", protocol.DecodeUpdate( buff.Bytes())
				fmt.Println(protocol.DecodeUpdate(buff.Bytes()))
			} else {
				printing <- fmt.Sprintf("", err)
				return
			}
			//wazzup(err)
			//conn.Write([]byte(":P"))
			//conn.Write(byte[](":P"))
		}
		time.Sleep(time.Minute)
	}

	update = func(conn net.Conn) {
		time.Sleep(time.Second)
		fmt.Println("##########################")
		conn.Write(mkUpdate())

	}
	return
}

func main() {

	talk, update, _ := construct()
	go server(talk, ":9000")
	go server(update, ":9001")

	time.Sleep(time.Hour)
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
