package main

import (
	"../protocol"
	"bytes"
	"encoding/binary"

	"fmt"
	"net"
	"os"
	"sync"
	"time"
	//"time"
)

func main() {
	var addr string
	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = "127.0.0.1"
	}
	talkConn, err1 := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, protocol.MAIN_PORT))
	updateConn, err2 := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, protocol.UPDATE_PORT))
	wazzup(err1)
	wazzup(err2)

	var firstUpdate sync.WaitGroup
	firstUpdate.Add(1)

	var lang protocol.Lang
	var lockLang sync.Mutex
	var add string = "waiting..."
	var lockAdd sync.Mutex
	var id func() int32
	safePrintAdd := func() {
		lockAdd.Lock()
		fmt.Println(add)
		lockAdd.Unlock()
	}

	go func(conn net.Conn) {
		first := true

		func() {
			arr := make([]byte, 4)
			_, err := conn.Read(arr)
			if err != nil {
				fmt.Println(err)
			}
			buf := bytes.NewBuffer(make([]byte, 0, 4))
			buf.Write(arr)
			var tempid int32
			binary.Read(buf, binary.BigEndian, &tempid)
			id = func() int32 {
				return tempid
			}
		}()

		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				continue
			}
			lockAdd.Lock()
			lockLang.Lock()
			add, lang = protocol.DecodeUpdate(buff[:n])
			lockAdd.Unlock()
			lockLang.Unlock()
			if first {
				first = false
				firstUpdate.Done()
			}

		}
		conn.Close()
	}(updateConn)
	firstUpdate.Wait()

	func() {
		buf := bytes.NewBuffer(make([]byte, 0, 4))
		binary.Write(buf, binary.BigEndian, id())
		buf.Bytes()
		talkConn.Write(buf.Bytes())
	}()
	time.Sleep(time.Millisecond * 1000)

	for {

		buff := make([]byte, 10)
		n, _ := talkConn.Read(buff)

		ustext, clear, rek, _ := protocol.DecodePrintCMD(buff[:n], lang)

		if clear {

			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			safePrintAdd()
			fmt.Println("\n\n\n")
		}
		fmt.Println(ustext)
		switch rek {
		case protocol.SEND_UINT:
			fmt.Println("\n\n\n")
			var e uint64
			fmt.Print("> ")
			fmt.Scanln(&e)
			talkConn.Write(protocol.MakeSendUintMSG(e))
			break
		case protocol.SEND_INT:
		}
	}
}

func wazzup(err error) {
	if err != nil {
		panic(err)
	}
}
