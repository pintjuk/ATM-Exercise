package main

import (
	"fmt"
	"net"
	"../Protocol"
)

func constructQuery() (query func(net.Conn), rconsol <-chan string) {
	consol := make(chan string, 5)
	rconsol = consol
	query = func(conn net.Conn) {
		consol <- "hej"
		conn.Write(byte[](":P"))
		conn.Close()
		close(consol)
	}
	return
}

func main() {
	query, consol := constructQuery()
	go server(query, ":9001")

	for i := range consol {
		fmt.Println(i)
	}
}
