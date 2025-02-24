package main 
import (
	"fmt"
	"net"
	"bufio"   //用来实现数据流的功能
	"os"
)

func main(){
	//这个方法是用来连接服务端的，发送请求  
	conn,err := net.Dial("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("连接出错了")    
		return
	}
	for {
    //现在准备实现发送单行数据
	reader := bufio.NewReader(os.Stdin)  //os.Stdin表示标准输入——代之终端  
	//读取一行输入，发送给服务端  
	line,err := reader.ReadString('\n')  //表示按照\n为隔断符号  
	//当然也可以使line = strings.Trim(line,"\n\r")表示空格和换行去掉,需要引入string 包   
	if line == "exit\n" {
		return 
	} 
	if err != nil {
		fmt.Println("reader error")
	}

	n,err := conn.Write([]byte(line))  //	这个方法可能会在超过某个固定时间限制后返回错误，里面需要byte类型的切片
	//会返回一个字节长度  
	if err != nil{
		fmt.Println("con.Write err=",err)	
	}
	fmt.Printf("客户端发送了 %v 个字节\n",n )

}
}