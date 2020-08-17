package main

import (
	"bufio"
	"fmt"
	"net"
)

type ServerRequestHandler struct {
	Port    string
	ConnTCP net.Conn
}

func (srh *ServerRequestHandler) Listen() (err error) {
	fmt.Println("Listen 1")

	ln, err := net.Listen("tcp", ":"+srh.Port)
	fmt.Println("Listen 2")

	fmt.Println(srh)
	fmt.Println(ln)
	fmt.Println(err)
	fmt.Println("Listen 3")

	if err != nil {
		return err
	}
	fmt.Println("Listen 4")

	conn, err2 := ln.Accept()
	fmt.Println("Listen 5")

	fmt.Println(conn)
	fmt.Println(err2)
	fmt.Println("Listen 6")

	if err2 == nil {
		srh.ConnTCP = conn
		fmt.Println("Listen 7")

		return nil
	} else {
		fmt.Println("Listen 8")

		return err2
	}
}

func (srh *ServerRequestHandler) Send(message []byte) error {
	_, err := srh.ConnTCP.Write(append(message, '\n'))
	return err
}

func (srh *ServerRequestHandler) Receive() ([]byte, error) {
	message, err := bufio.NewReader(srh.ConnTCP).ReadBytes('\n')
	return message, err
}
