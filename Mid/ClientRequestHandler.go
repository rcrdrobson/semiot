package main

import (
	"bufio"
	"fmt"
	"net"
)

type ClientRequestHandler struct {
	Host    string
	Port    string
	ConnTCP net.Conn
}

func (crh *ClientRequestHandler) Connect() (err error) {
	conn, err := net.Dial("tcp", crh.Host+":"+crh.Port)
	if err == nil {
		crh.ConnTCP = conn
		return nil
	} else {
		return err
	}
}

func (crh *ClientRequestHandler) Send(message []byte) error {
	_, err := fmt.Fprintf(crh.ConnTCP, string(message)+"\n")
	return err
}

func (crh *ClientRequestHandler) Receive() ([]byte, error) {
	message, err := bufio.NewReader(crh.ConnTCP).ReadBytes('\n')
	return message, err
}

func (crh *ClientRequestHandler) Close() error {
	err := crh.ConnTCP.Close()
	return err
}
