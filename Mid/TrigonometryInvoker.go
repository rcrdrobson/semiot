package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type TrigonometryInvoker struct{}

func (t *TrigonometryInvoker) Invoke(trigonometryProxy TrigonometryProxy) {
	srh := new(ServerRequestHandler)
	srh.Port = trigonometryProxy.Port

	fmt.Println(trigonometryProxy)
	fmt.Println(srh)
	var msgToBeUnmarshalled []byte
	var msgMarshalled []byte

	marshaller := new(Marshaller)
	terminator := new(Terminator)

	for {
		fmt.Println("loop")
		srh.Listen()
		msgToBeUnmarshalled, _ = srh.Receive()
		fmt.Println("msgToBeUnmarshalled")
		fmt.Println(msgToBeUnmarshalled)

		msgUnmarshalled := marshaller.Unmarshal(msgToBeUnmarshalled)
		fmt.Println("msgUnmarshalled")
		fmt.Println(msgUnmarshalled)
		op := msgUnmarshalled.MessageBody.RequestHeader.Operation
		fmt.Println("op")
		fmt.Println(op)
		switch op {
		case "Sin":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])
			fmt.Println("x")
			fmt.Println(x)
			jsonBody := "{\"x\":" + x + "}"
			fmt.Println("jsonBodyop")
			fmt.Println(jsonBody)
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", trigonometryProxy.HostRunner, trigonometryProxy.PortRunner, "trigonometry/trigonometry", op), body)
			fmt.Println("req")
			fmt.Println(req)
			fmt.Println("err")
			fmt.Println(err)
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
			fmt.Println("strResult")
			fmt.Println(strResult)
			terminator.Result = strResult

			fmt.Println("terminator")
			fmt.Println(terminator)
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
			fmt.Println("pubMessageToBeMarshalled")
			fmt.Println(pubMessageToBeMarshalled)
			msgMarshalled = marshaller.Marshal(pubMessageToBeMarshalled)
			fmt.Println("msgMarshalled")
			fmt.Println(msgMarshalled)
			srh.Send(msgMarshalled)
		case "Cos":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])

			jsonBody := "{\"x\":" + x + "}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", trigonometryProxy.HostRunner, trigonometryProxy.PortRunner, "trigonometry/trigonometry", op), body)

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
		case "Tan":
			x := string(msgUnmarshalled.MessageBody.RequestBody.Parameters[0])

			jsonBody := "{\"x\":" + x + "}"
			var body io.Reader
			body = strings.NewReader(jsonBody)

			req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/call/%v/%v", trigonometryProxy.HostRunner, trigonometryProxy.PortRunner, "trigonometry/trigonometry", op), body)

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
			fmt.Println("send")
			srh.Send(msgMarshalled)
		}
	}
}
