package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, HTTPS!")
}

func main() {
	mux := http.NewServeMux()  //创建一个 新的 HTTP 多路复用器
	mux.HandleFunc("/hello", handler)  //mux 变量存储了这个 路由管理器

	server := &http.Server{
		Addr:    "127.0.0.1:8090",    //":8090" 表示监听本机所有网卡上的 8090 端口 这里表示只监听本地 
		Handler: mux,        //表示的就是采用这个路由器  
	}

	log.Println("Starting HTTPS server on https://localhost:8090") //日志文件记录
	err := server.ListenAndServeTLS("server.crt", "server.key")  //监听和使用方法		
	if err != nil {
		log.Fatal("ListenAndServeTLS error:", err)
	}
}
