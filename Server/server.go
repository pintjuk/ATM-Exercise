package main

import (
	//"../protocol"
	"fmt"
	"net"
	"strings"
)

func construct() (talk func(net.Conn), update func(net.Conn), rprinting <-chan string) {
	printing := make(chan string, 5)
	rprinting = printing

	talk = func(conn net.Conn) {

		printing <- strings.Join([]string{"new connection from:", conn.RemoteAddr().String()}, " ")
		//conn.Write(byte[](":P"))
		close(printing)
	}

	update = func(conn net.Conn) {
		conn.Write([]byte(":P"))
	}
	return
}

func main() {

	talk, update, printing := construct()
	go server(talk, ":9000")
	go server(update, ":9001")

	for i := range printing {
		fmt.Println(i)
	}
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
