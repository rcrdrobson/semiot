package main

type Invoker struct {
	OperationName string
	Host          string
	Port          string
	Parameters    [][]byte
}
