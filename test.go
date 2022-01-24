package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if nil != err {
		log.Println(err)
	}

	go func() {
		data := make([]byte, 4096)

		for {
			n, err := conn.Read(data)
			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Server send : " + string(data[:n]))
		}
	}()

	for {
		var s string
		fmt.Scanln(&s)
		conn.Write([]byte(s))
	}
}
