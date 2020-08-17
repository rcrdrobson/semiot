package main

import (
	"encoding/json"
)

type Marshaller struct{}

func (m *Marshaller) Marshal(message interface{}) []byte {
	marshalMessage, _ := json.Marshal(message)
	return marshalMessage
}

func (m *Marshaller) Unmarshal(message []byte) Message {
	var resultUnmarshal Message
	json.Unmarshal(message, &resultUnmarshal)
	return resultUnmarshal
}
