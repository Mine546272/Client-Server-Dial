package main

import (
	"fmt"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", "localhost:3355")

	if err != nil {
		return
	}
	defer dial.Close()

	payload := "Dial is Successful"
	_, err = dial.Write([]byte(payload))
	if err != nil {
		fmt.Println(err)
		return
	}
}
