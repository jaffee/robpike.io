package main

import (
	"net"
	"time"
	"fmt"
)

func handleConn(conn net.Conn) {
	fmt.Printf("New Connection %v\n", conn.RemoteAddr())
	packet := "HTTP/1.1 200 OK\nContent-Type: text/plain; charset=ascii\n\n"
	buff := make([]byte, len(packet))
	copy(buff[:], packet)
	_, err := conn.Write(buff)
	if err != nil {
		fmt.Printf("Lost Connection %v\n", conn.RemoteAddr())
		return
	}
	tb := []byte{74, 74}
	for {
//		t := time.Now()
		_, err = conn.Write(tb)
//		elapsed := time.Since(t)
//		fmt.Printf("Elapsed time for write: %v\n", elapsed)
		if err != nil {
			fmt.Printf("Lost Connection %v\n", conn.RemoteAddr())
			return
		}
		time.Sleep(time.Second)
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
