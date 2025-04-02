package main 
import (
	"log"
	"net"
	"bufio"
	"os"
)

func main(){
	serverAddr := ":18889"
	conn ,err := net.Dial("tcp",serverAddr)
	if err != nil {
		log.Fatalf("客户端连接失败")  
	}
	defer conn.Close()
	reader  := bufio.NewReader(os.Stdin)
	//返回一个默认缓冲区的reader	
	for{
		log.Println("输入发送的请求")
		userInput , _ :=reader.ReadString('\n')
		// 遇到换行符号自动发送
		_,err =conn.Write([]byte(userInput))
		if err !=nil{
			log.Println("发送失败%v",err)
		}
	
	}
}