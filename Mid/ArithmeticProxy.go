package main

import "strconv"

// Information Registered in the Name Service by the Object Server
type ArithmeticProxy struct {
	Host       string
	Port       string
	HostRunner string
	PortRunner string
}

func (ap *ArithmeticProxy) Sum(x, y float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	// Receives information from the Client
	marshalX := marshaller.Marshal(x)
	invoker.Parameters = append(invoker.Parameters, marshalX)

	marshalY := marshaller.Marshal(y)
	invoker.Parameters = append(invoker.Parameters, marshalY)

	// Prepares the invocation to the Requestor
	invoker.Host = ap.Host
	invoker.Port = ap.Port
	invoker.OperationName = "Sum"

	// Invokes the Requestor (synchronous)
	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	// Sends reply to Client
	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}

func (ap *ArithmeticProxy) Sub(x, y float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	marshalX := marshaller.Marshal(x)
	invoker.Parameters = append(invoker.Parameters, marshalX)

	marshalY := marshaller.Marshal(y)
	invoker.Parameters = append(invoker.Parameters, marshalY)

	invoker.Host = ap.Host
	invoker.Port = ap.Port
	invoker.OperationName = "Sub"

	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}

func (ap *ArithmeticProxy) Mul(x, y float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	marshalX := marshaller.Marshal(x)
	invoker.Parameters = append(invoker.Parameters, marshalX)

	marshalY := marshaller.Marshal(y)
	invoker.Parameters = append(invoker.Parameters, marshalY)

	invoker.Host = ap.Host
	invoker.Port = ap.Port
	invoker.OperationName = "Mul"

	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}

func (ap *ArithmeticProxy) Div(x, y float64) float64 {
	invoker := new(Invoker)
	marshaller := new(Marshaller)

	marshalX := marshaller.Marshal(x)
	invoker.Parameters = append(invoker.Parameters, marshalX)

	marshalY := marshaller.Marshal(y)
	invoker.Parameters = append(invoker.Parameters, marshalY)

	invoker.Host = ap.Host
	invoker.Port = ap.Port
	invoker.OperationName = "Div"

	requestor := new(Requestor)
	terminator := requestor.Invoke(*invoker)

	result := terminator.Result
	resultUnmarshal := marshaller.Unmarshal(result)
	opResultBytes := resultUnmarshal.MessageBody.ReplyBody.OperationResult
	opResultStr := string(opResultBytes)
	f, _ := strconv.ParseFloat(opResultStr, 64)
	return f
}
