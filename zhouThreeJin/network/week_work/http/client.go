/*
tr := &http.Transport{
	MaxIdleConns:       10,
	IdleConnTimeout:    30 * time.Second,
	DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
//相当于生成一个结构体——里面定义了HTTP头部详细信息  


相当于先定义了这个函数，如果实体后会用到了就会触发这个
func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	// 限制最多只能重定向 3 次
	if len(via) >= 3 {
		return fmt.Errorf("stopped after 3 redirects")
	}
	return nil // 允许重定向
}  

func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,  //引用了这个函数
	}
 

*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8090/hello"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body) 
//ioutil.ReadAll(resp.Body) 读取 resp.Body 中的所有数据并返回 []byte（字节切片）
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response from server:", string(body))
}
