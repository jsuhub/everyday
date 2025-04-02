package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 读取自签证书
	
	caCert, err := os.ReadFile("server.crt")   //客户端读取证书
	if err != nil {
		log.Fatal("无法读取证书:", err)
	}

	// 只是在当前 Go 进程中创建了一个 CA 证书池（Cert Pool），并添加了 server.crt 证书。
	//这只影响当前 Go 客户端的 TLS 连接，并不会改变系统的全局环境变量或 CA 证书存储。
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// 因为修改了连接池所以修改了TLS的配置项，所以重新生成的实力类别
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: caCertPool},
	}
	
	client := &http.Client{Transport: tr}  

	// 发送 HTTPS 请求
	resp, err := client.Get("https://localhost:8090/hello")
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}
