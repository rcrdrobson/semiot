package main

type RequestHeader struct {
	Context          string
	RequestId        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}
