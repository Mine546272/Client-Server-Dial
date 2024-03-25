package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3355")
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}
		go func(c net.Conn) {
			fmt.Println("Accepted")
			buf := make([]byte, 100)
			_, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("%s", buf)
		}(conn)

	}

}
