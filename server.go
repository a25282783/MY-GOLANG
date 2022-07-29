package main

import (
	"fmt"
	"net"
)

type Server struct {
	ip   string
	port int
}

func NewServer(ip string, port int) *Server {
	server := Server{
		ip:   ip,
		port: port,
	}

	return &server
}

func (this *Server) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.ip, this.port))
	if err != nil {
		fmt.Print(err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}

		go func(conn net.Conn) {
			fmt.Print("連線成功...")
		}(conn)

	}
}
