package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "8005")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		_, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println("err")
		} else {
			fmt.Println("ok")
		}

	}
}
