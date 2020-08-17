package main

import (
	"fmt"
	"strconv"
)

// Information Registered in the Name Service by the Object Server
type TrigonometryProxy struct {
	Host       string
	Port       string
	HostRunner string
	PortRunner string
}

func (tp *TrigonometryProxy) Sin(angle float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	// Receives information from the Client
	marshalAngle := marshaller.Marshal(angle)
	invoker.Parameters = append(invoker.Parameters, marshalAngle)

	// Prepares the invocation to the Requestor
	invoker.Host = tp.Host
	invoker.Port = tp.Port
	invoker.OperationName = "Sin"

	// Invokes the Requestor (synchronous)
	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	// Sends reply to Client
	result := terminator.Result
	fmt.Println("Result")
	fmt.Println(result)

	resultUnmarshal := marshaller.Unmarshal(result)
	fmt.Println("resultUnmarshal")
	fmt.Println(resultUnmarshal)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	//opResultStr := string(opResultBytes)
	fmt.Println("opResultBytes")
	fmt.Println(opResultBytes)
	f, _ := strconv.ParseFloat(string(result), 64)
	return f
}

func (tp *TrigonometryProxy) Cos(angle float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	marshalAngle := marshaller.Marshal(angle)
	invoker.Parameters = append(invoker.Parameters, marshalAngle)

	invoker.Host = tp.Host
	invoker.Port = tp.Port
	invoker.OperationName = "Cos"

	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}

func (tp *TrigonometryProxy) Tan(angle float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	marshalAngle := marshaller.Marshal(angle)
	invoker.Parameters = append(invoker.Parameters, marshalAngle)

	invoker.Host = tp.Host
	invoker.Port = tp.Port
	invoker.OperationName = "Tan"

	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}
