package main

type Requestor struct {
}

func (r *Requestor) Invoke(inv Invoker) *Terminator {
	crh := new(ClientRequestHandler)
	crh.Host = inv.Host
	crh.Port = inv.Port

	marshaller := new(Marshaller)
	terminator := new(Terminator)

	var msgMarshalled []byte
	var msgToBeUnmarshalled []byte

	// Creates the message to be transmitted
	requestHeader := new(RequestHeader)
	requestHeader.Operation = inv.OperationName
	requestHeader.ResponseExpected = true
	requestHeader.Context = ""
	requestHeader.RequestId = 0
	requestHeader.ObjectKey = 0

	requestBody := new(RequestBody)
	requestBody.Parameters = inv.Parameters

	// Default values - not useds
	messageHeader := new(MessageHeader)
	messageHeader.Magic = "MIOP"
	messageHeader.ByteOrder = false
	messageHeader.MessageSize = 0
	messageHeader.MessageType = 0
	messageHeader.Version = 0

	messageBody := new(MessageBody)
	messageBody.RequestBody = *requestBody
	messageBody.RequestHeader = *requestHeader

	message := new(Message)
	message.MessageBody = *messageBody
	message.MessageHeader = *messageHeader

	// Serializes the message request
	msgMarshalled = marshaller.Marshal(message)

	crh.Connect()

	// Send the message
	crh.Send(msgMarshalled)

	// Reveice the message
	msgToBeUnmarshalled, _ = crh.Receive()

	crh.Close()

	// Deserializes the message
	msgUnmarshalled := marshaller.Unmarshal(msgToBeUnmarshalled)

	// Send response to Object Proxy
	terminator.Result = msgUnmarshalled.MessageBody.ReplyBody.OperationResult
	return terminator
}
