package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ArithmeticInvoker struct{}

func (a *ArithmeticInvoker) Invoke(arithmeticProxy ArithmeticProxy) {
	srh := new(ServerRequestHandler)
	srh.Port = arithmeticProxy.Port
	srh.Listen()

	var msgToBeUnmarshalled []byte
	var msgMarshalled []byte

	marshaller := new(Marshaller)
	terminator := new(Terminator)

	for {
		msgToBeUnmarshalled, _ = srh.Receive()
		msgUnmarshalled := marshaller.Unmarshal(msgToBeUnmarshalled)

		op := msgUnmarshalled.MessageBody.RequestHeader.Operation

		switch op {
		case "Sum":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])
			y := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[1])

			jsonBody := "{\"x\":\"" + x + "\",\"hyost\": \"" + y + "\"}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", arithmeticProxy.HostRunner, arithmeticProxy.PortRunner, "arithmetic/arithmetic", op), body)
			var res *http.Response

			for i := 0; i < 2000; i++ {
				res, err = http.DefaultClient.Do(req)
				fmt.Println(err)
				if err == nil {
					fmt.Printf("Connection with '%s' is OK", op)
					break
				}
				time.Sleep(10 * time.Millisecond)
			}

			strResult, _ := ioutil.ReadAll(res.Body)
			terminator.Result = strResult

			replyHeader := new(ReplyHeader)
			replyHeader.ReplyStatus = 0
			replyHeader.RequestId = 0
			replyHeader.ServiceContext = ""

			replyBody := new(ReplyBody)
			replyBody.OperationResult = terminator.Result

			messageHeader := new(MessageHeader)
			messageHeader.Magic = "protocolo"
			messageHeader.ByteOrder = false
			messageHeader.MessageSize = 0
			messageHeader.MessageType = 0
			messageHeader.Version = 0

			messageBody := new(MessageBody)
			messageBody.ReplyBody = *replyBody
			messageBody.ReplyHeader = *replyHeader

			pubMessageToBeMarshalled := new(Message)
			pubMessageToBeMarshalled.MessageBody = *messageBody
			pubMessageToBeMarshalled.MessageHeader = *messageHeader
			msgMarshalled = marshaller.Marshal(pubMessageToBeMarshalled)

			srh.Send(msgMarshalled)
		case "Sub":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])
			y := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[1])

			jsonBody := "{\"x\":\"" + x + "\",\"hyost\": \"" + y + "\"}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", arithmeticProxy.HostRunner, arithmeticProxy.PortRunner, "arithmetic/arithmetic", op), body)

			var res *http.Response

			for i := 0; i < 2000; i++ {
				res, err = http.DefaultClient.Do(req)
				fmt.Println(err)
				if err == nil {
					fmt.Printf("Connection with '%s' is OK", op)
					break
				}
				time.Sleep(10 * time.Millisecond)
			}

			strResult, _ := ioutil.ReadAll(res.Body)
			terminator.Result = strResult

			replyHeader := new(ReplyHeader)
			replyHeader.ReplyStatus = 0
			replyHeader.RequestId = 0
			replyHeader.ServiceContext = ""

			replyBody := new(ReplyBody)
			replyBody.OperationResult = terminator.Result

			messageHeader := new(MessageHeader)
			messageHeader.Magic = "protocolo"
			messageHeader.ByteOrder = false
			messageHeader.MessageSize = 0
			messageHeader.MessageType = 0
			messageHeader.Version = 0

			messageBody := new(MessageBody)
			messageBody.ReplyBody = *replyBody
			messageBody.ReplyHeader = *replyHeader

			pubMessageToBeMarshalled := new(Message)
			pubMessageToBeMarshalled.MessageBody = *messageBody
			pubMessageToBeMarshalled.MessageHeader = *messageHeader
			msgMarshalled = marshaller.Marshal(pubMessageToBeMarshalled)

			srh.Send(msgMarshalled)
		case "Mul":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])
			y := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[1])

			jsonBody := "{\"x\":\"" + x + "\",\"hyost\": \"" + y + "\"}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", arithmeticProxy.HostRunner, arithmeticProxy.PortRunner, "arithmetic/arithmetic", op), body)

			var res *http.Response

			for i := 0; i < 2000; i++ {
				res, err = http.DefaultClient.Do(req)
				fmt.Println(err)
				if err == nil {
					fmt.Printf("Connection with '%s' is OK", op)
					break
				}
				time.Sleep(10 * time.Millisecond)
			}
			strResult, _ := ioutil.ReadAll(res.Body)
			terminator.Result = strResult

			replyHeader := new(ReplyHeader)
			replyHeader.ReplyStatus = 0
			replyHeader.RequestId = 0
			replyHeader.ServiceContext = ""

			replyBody := new(ReplyBody)
			replyBody.OperationResult = terminator.Result

			messageHeader := new(MessageHeader)
			messageHeader.Magic = "protocolo"
			messageHeader.ByteOrder = false
			messageHeader.MessageSize = 0
			messageHeader.MessageType = 0
			messageHeader.Version = 0

			messageBody := new(MessageBody)
			messageBody.ReplyBody = *replyBody
			messageBody.ReplyHeader = *replyHeader

			pubMessageToBeMarshalled := new(Message)
			pubMessageToBeMarshalled.MessageBody = *messageBody
			pubMessageToBeMarshalled.MessageHeader = *messageHeader
			msgMarshalled = marshaller.Marshal(pubMessageToBeMarshalled)

			srh.Send(msgMarshalled)
		case "Div":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])
			y := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[1])

			jsonBody := "{\"x\":\"" + x + "\",\"hyost\": \"" + y + "\"}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", arithmeticProxy.HostRunner, arithmeticProxy.PortRunner, "arithmetic/arithmetic", op), body)

			var res *http.Response

			for i := 0; i < 2000; i++ {
				res, err = http.DefaultClient.Do(req)
				fmt.Println(err)
				if err == nil {
					fmt.Printf("Connection with '%s' is OK", op)
					break
				}
				time.Sleep(10 * time.Millisecond)
			}

			strResult, _ := ioutil.ReadAll(res.Body)
			terminator.Result = strResult

			replyHeader := new(ReplyHeader)
			replyHeader.ReplyStatus = 0
			replyHeader.RequestId = 0
			replyHeader.ServiceContext = ""

			replyBody := new(ReplyBody)
			replyBody.OperationResult = terminator.Result

			messageHeader := new(MessageHeader)
			messageHeader.Magic = "protocolo"
			messageHeader.ByteOrder = false
			messageHeader.MessageSize = 0
			messageHeader.MessageType = 0
			messageHeader.Version = 0

			messageBody := new(MessageBody)
			messageBody.ReplyBody = *replyBody
			messageBody.ReplyHeader = *replyHeader

			pubMessageToBeMarshalled := new(Message)
			pubMessageToBeMarshalled.MessageBody = *messageBody
			pubMessageToBeMarshalled.MessageHeader = *messageHeader
			msgMarshalled = marshaller.Marshal(pubMessageToBeMarshalled)

			srh.Send(msgMarshalled)
		}

	}
}
