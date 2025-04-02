package SS5 
import (
	"fmt"
	"log"
	"net"
)

func StartSocks5_server(config *SSConfig){
	listener ,err := net.Listen("tcp",fmt.Sprintf("127.0.0.1:%d",config.LocalPort))
	//设置一个长期的监听本地端口的监听器 
		//格式化的数据转换为字符串，实现监听tcp 并且为LocalPort	
	if err!= nil {
		log.Fatalf("服务器监听本地端口失败%v",err)
		//是一个终止语句  
	}
	log.Println("服务器监听本地端口",config.LocalPort)  

	for {
		conn ,err := listener.Accept()  
		//Accept 是会阻塞的，即没有新的连接的时候就会一直停止在这里  
		if err!=nil{ 
			log.Fatalf("服务器连接已经关闭了%v",err)
			continue
		}
		log.Println("服务器已经监听到数据传来 ",config.LocalPort)  
		go handleServer(conn ,config)
				//异步请求，从而并行处理生成多条线程  
	}
}

func handleServer(conn net.Conn , config *SSConfig){
	log.Println("服务器正处理") 
	receirveTraffic(conn,config)
}


func receirveTraffic(conn net.Conn,config *SSConfig){
	//处理数据并转发
	log.Println("数据开始接受")
	defer conn.Close()
	buf := make([]byte,4096) 
	for {
		n,err := conn.Read(buf)
		//Read() 方法会尝试从 conn 读取数据，并将其存储到 （buf）中
		if err !=nil {
			log.Fatalf("出现错误%v",err)
			//表示数据读取完毕了
			break
		}
		log.Println(" 从客户端读取数据成功")
		decryptrData,err := DecryptAES256GCM(buf[:n],[]byte(config.Password))
		if err !=nil {
			log.Fatalf("出现错误%v",err)
		}
		
		log.Println(" 数据正在解密")
		//进行加密 
		log.Println("数据解密成功为：",decryptrData)
	}
}

