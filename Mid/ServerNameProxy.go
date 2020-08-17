package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type ServerNameProxy struct {
	Host         string
	Port         string
	TypeRegistry map[string]reflect.Type
}

func (sn *ServerNameProxy) Init() {
	sn.TypeRegistry = make(map[string]reflect.Type)
	myTypes := []interface{}{ArithmeticProxy{}, TrigonometryProxy{}}
	for _, v := range myTypes {
		sn.TypeRegistry[fmt.Sprintf("%T", v)] = reflect.TypeOf(v)
		sn.makeInstance(fmt.Sprintf("%T", v))
	}
}

func (sn *ServerNameProxy) makeInstance(name string) interface{} {
	v := reflect.New(sn.TypeRegistry[name]).Elem()
	// Maybe fill in fields here if necessary
	return v.Interface()
}

func (sn *ServerNameProxy) LookupArithmetic(name string) ArithmeticProxy {
	jsonBody := "{\"name\":\"" + name + "\"}"
	var body io.Reader
	body = strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/lookup/", sn.Host, sn.Port), body)

	var res *http.Response

	for i := 0; i < 2000; i++ {
		res, err = http.DefaultClient.Do(req)
		fmt.Println(err)
		if err == nil {
			fmt.Printf("Connection with '%s' is OK", name)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	message, err := ioutil.ReadAll(res.Body)
	messageReader := bytes.NewReader(message)
	hostRunner, portRunner, port := extractLookup(messageReader)

	var arithmeticProxy ArithmeticProxy
	arithmeticProxy.HostRunner = hostRunner
	arithmeticProxy.PortRunner = portRunner
	arithmeticProxy.Port = port
	return arithmeticProxy
}

func (sn *ServerNameProxy) LookupTrigonometry(name string) TrigonometryProxy {
	jsonBody := "{\"name\":\"" + name + "\"}"
	var body io.Reader
	body = strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/lookup/", sn.Host, sn.Port), body)

	var res *http.Response

	for i := 0; i < 2000; i++ {
		res, err = http.DefaultClient.Do(req)
		fmt.Println(err)
		if err == nil {
			fmt.Printf("Connection with '%s' is OK", name)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	message, err := ioutil.ReadAll(res.Body)
	messageReader := bytes.NewReader(message)
	hostRunner, portRunner, port := extractLookup(messageReader)

	var trigonometryProxy TrigonometryProxy
	trigonometryProxy.HostRunner = hostRunner
	trigonometryProxy.PortRunner = portRunner
	trigonometryProxy.Port = port

	return trigonometryProxy
}

func extractLookup(jsonBodyReq io.Reader) (hostRunner, portRunner, port string) {
	var jsonBody interface{}
	err := json.NewDecoder(jsonBodyReq).Decode(&jsonBody)
	if err != nil {
		return
	}

	var bodyData = jsonBody.(map[string]interface{})
	return bodyData["hostRunner"].(string), bodyData["portRunner"].(string), bodyData["port"].(string)
}

func (sn *ServerNameProxy) List(name string) (messageRes string, e error) {
	jsonBody := "{\"name\":\"" + name + "\"}"
	var body io.Reader
	body = strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/list/", sn.Host, sn.Port), body)

	var res *http.Response

	for i := 0; i < 2000; i++ {
		res, err = http.DefaultClient.Do(req)
		fmt.Println(err)
		if err == nil {
			fmt.Printf("Connection with '%s' is OK", name)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	message, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return string(message), nil
	} else {
		return "", err
	}
}

func (sn *ServerNameProxy) Bind(name, host, port string) (messageRes string, e error) {
	jsonBody := "{\"name\":\"" + name + "\",\"host\": \"" + host + "\",\"port\": \"" + port + "\"}"
	var body io.Reader
	body = strings.NewReader(jsonBody)
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:%v/bind/", sn.Host, sn.Port), body)

	var res *http.Response

	for i := 0; i < 2000; i++ {
		res, err = http.DefaultClient.Do(req)
		//fmt.Println(err)
		if err == nil {
			fmt.Printf("Connection with '%s' is OK", name)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	message, err := ioutil.ReadAll(res.Body)
	if err == nil {
		return string(message), nil
	} else {
		return "", err
	}
}
