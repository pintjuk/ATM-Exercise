package main

import (
	"../protocol"
	//"bytes"
	//"bytes"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	var addr string
	if len(os.Args) > 1 {
		addr = os.Args[1]
	} else {
		addr = "127.0.0.1"
	}
	_, err1 := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, protocol.MAIN_PORT))
	UpdateConn, err2 := net.Dial("tcp", fmt.Sprintf("%s:%d", addr, protocol.UPDATE_PORT))
	wazzup(err1)
	wazzup(err2)

	var lang protocol.Lang
	var lockLang sync.Mutex
	var add string = "waiting..."
	var lockAdd sync.Mutex

	go func(conn net.Conn) {
		for {
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				//fmt.Println("Error", err)
				continue
			}
			lockAdd.Lock()
			lockLang.Lock()
			add, lang = protocol.DecodeUpdate(buff[:n])
			lockAdd.Unlock()
			lockLang.Unlock()
		}
		conn.Close()
	}(UpdateConn)

	for {
		time.Sleep(time.Second)
		lockAdd.Lock()
		lockLang.Lock()
		fmt.Println(add, lang)
		lockLang.Unlock()
		lockAdd.Unlock()
	}
}

func wazzup(err error) {
	if err != nil {
		panic(err)
	}
}
