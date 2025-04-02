/*
http.Handle("/foo", fooHandler)
表示的是注册路由


type FooHandler struct{}

func (f FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from /foo")
}

var fooHandler FooHandler  定义fooHandler 必须要满足的条件和返回值  

*/
package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你好") 
	//fmt.Fprintf 是 Go 语言 fmt 包中的一个函数，它的作用是 格式化字符串并写入到指定的 io.Writer。 
	
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port := ":8090"
	fmt.Println("Server is running on port", port)

	if err := http.ListenAndServe(port, nil);err != nil {
		fmt.Println("Error starting server:", err)
	}
}
