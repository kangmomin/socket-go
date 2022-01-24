package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	app, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer app.Close()

	for {
		conn, err := app.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		defer conn.Close()
		go ConnHandler(conn)
		go SendMsg(conn)
	}
}

func ConnHandler(conn net.Conn) {
	recvBuffur := make([]byte, 4096)
	for {
		n, err := conn.Read(recvBuffur)
		if err != nil {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}

		fmt.Println(string(recvBuffur[:n]))
	}
}

func SendMsg(conn net.Conn) {
	msg := ""
	fmt.Scan(&msg)

	conn.Write([]byte(msg))
}
