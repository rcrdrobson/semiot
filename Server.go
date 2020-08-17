package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	bindEndpoint   = "/bind/"
	lookupEndpoint = "/lookup/"
	listEndpoint   = "/list/"
	port           = ":9090"
)

var poolMap = make(map[string][]string)

func main() {
	fmt.Println("Starting ServerName at port '" + port + "'")
	http.HandleFunc(bindEndpoint, bind)
	http.HandleFunc(lookupEndpoint, lookup)
	http.HandleFunc(listEndpoint, list)

	fmt.Println("Started ServerName")

	http.ListenAndServe(port, nil)

}

func lookup(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Lookup func")
	name := extractLookup(res, req.Body)
	address := poolMap[name][0]
	addressSplitted := strings.Split(address, ":")
	fmt.Println(addressSplitted)

	// Shift on the poolMap
	poolMap[name] = poolMap[name][1:]
	poolMap[name] = append(poolMap[name], address)

	fmt.Println(addressSplitted)
	fmt.Println("0" + addressSplitted[0])
	fmt.Println("1" + addressSplitted[1])
	fmt.Println("2" + addressSplitted[2])

	body := "{\"hostRunner\":\"" + addressSplitted[0] + "\",\"portRunner\":\"" + addressSplitted[1] + "\",\"port\":\"" + addressSplitted[2] + "\"}"
	fmt.Println("Lookup response body: " + body)

	res.Write([]byte(body))
	res.WriteHeader(http.StatusCreated)
}

func list(res http.ResponseWriter, req *http.Request) {
	fmt.Println("List func")
	fmt.Println(poolMap)

	var keys string
	for k := range poolMap {
		keys = keys + ", " + k
	}

	body := fmt.Sprintf(keys[2:])

	fmt.Println("List response body: " + body)

	res.Write([]byte(body))
	res.WriteHeader(http.StatusCreated)
}

func bind(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Bind func")
	name, hostRunner, portRunner, port := extractBind(res, req.Body)

	poolMap[name] = append(poolMap[name], hostRunner+":"+portRunner+":"+port)
	fmt.Println("Bind func")
	body := "Serverless '" + name + "' has binded"
	fmt.Println("Bind response body: " + body)
	res.Write([]byte(body))
	res.WriteHeader(http.StatusCreated)
}

func extractLookup(res http.ResponseWriter, jsonBodyReq io.Reader) (name string) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["name"].(string)
}

func extractBind(res http.ResponseWriter, jsonBodyReq io.Reader) (name, hostRunner, portRunner, port string) {
	fmt.Println("jsonBodyReq")
	fmt.Println(jsonBodyReq)
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		http.Error(res, err.Error(), 400)
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	fmt.Println("Bind response body: ")
	fmt.Println(bodyData)

	return bodyData["name"].(string), bodyData["hostRunner"].(string), bodyData["portRunner"].(string), bodyData["port"].(string)
}
