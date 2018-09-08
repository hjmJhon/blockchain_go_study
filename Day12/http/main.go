package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	//getBaidu()
	http.HandleFunc("/", hello)
	http.HandleFunc("/user/login", intercepterHandrler(login))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("lister fail")
	}

}
func intercepterHandrler(handlerFunc func(writer http.ResponseWriter, request *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		if i := recover(); i != nil {
			fmt.Println("err", i)
		} else {
			fmt.Println("success, continue")
		}
		handlerFunc(writer, request)
	}
}
func getBaidu() {
	response, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("get fail")
	}
	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		fmt.Println("read body fail")
		return
	}
	fmt.Println(string(bytes))
}

func login(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("login"))

}
func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("HandleFunc hello: hello")
	writer.Write([]byte("hello"))
}
