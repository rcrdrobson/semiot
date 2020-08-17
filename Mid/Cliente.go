package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//Connect with server name
	serverName := new(ServerNameProxy)
	serverName.Host = "localhost"
	serverName.Port = "9090"

	var t TrigonometryProxy
	t = serverName.LookupTrigonometry("trigonometry")
	iteracoes := 40
	t1 := time.Now()

	for i := 0; i < iteracoes; i++ {
		result := t.Sin(float64(1))
		fmt.Println("Result " + strconv.Itoa(i) + ":\t" + fmt.Sprintf("%f", result))
	}
	t2 := time.Now()
	x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
	fmt.Println(x)
}
