package main

import (
	"fmt"
	"net"
	"time"
)

func handleConn(conn net.Conn) {
	fmt.Printf("New Connection %v\n", conn.RemoteAddr())
	packet := "HTTP/1.1 200 OK\nContent-Type: text/plain; charset=ascii\nTransfer-Encoding: chunked\n\n"
	buff := make([]byte, len(packet))
	copy(buff[:], packet)
	_, err := conn.Write(buff)
	if err != nil {
		fmt.Printf("Lost Connection %v\n", conn.RemoteAddr())
		return
	}
	tb := []byte{74, 74}
	for {
		_, err = conn.Write(tb)
		if err != nil {
			fmt.Printf("Lost Connection %v\n", conn.RemoteAddr())
			return
		}
		time.Sleep(time.Second / 2)
	}
}

func main() {
	l, err := net.Listen("tcp", ":5566")
	if err != nil {
		return
	}
	for {
		conn, err := l.Accept()
		if err == nil {
			go handleConn(conn)
		}
	}
}
