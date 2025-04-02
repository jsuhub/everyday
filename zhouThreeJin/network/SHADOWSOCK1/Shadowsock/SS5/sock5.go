package SS5
import (
	"fmt"
	"log"
	"net"
	"errors"
)

func StartSocks5(config *SSConfig){
	listener ,err := net.Listen("tcp",fmt.Sprintf("127.0.0.1:%d",config.LocalPort))
	//设置一个长期的监听本地端口的监听器 
		//格式化的数据转换为字符串，实现监听tcp 并且为LocalPort	
	if err!= nil {
		log.Fatalf("无法启动本地代理%v",err)
		//是一个终止语句  
	}
	log.Println("SOCKS5 代理成功监听",config.LocalPort)  

	for {
		conn ,err := listener.Accept()  
		//Accept 是会阻塞的，即没有新的连接的时候就会一直停止在这里  
		if err!=nil{ 
			log.Fatalf("连接已经关闭了%v",err)
			continue
		}
		log.Println("SOCKS5 正在持续监听 ",config.LocalPort)  
		go handleClient(conn ,config)
				//异步请求，从而并行处理生成多条线程  
	}
}

	func handleClient(conn net.Conn,config *SSConfig){
		forwardTraffic(conn,config)
	}

	func connectToRemote(config *SSConfig)(net.Conn ,error ){
		serverAddr := fmt.Sprintf("%s:%d",config.Server,config.ServerPort) //把请求发送到服务端哪里去
		conn ,err := net.Dial("tcp",serverAddr)   
		//当使用 TCP 时，如果主机解析为多个 IP 地址，Dial 将按顺序尝试每个 IP 地址，直到成功为止。
		if err != nil {
			return nil,errors.New("error")
		}
		return conn,nil 
	}

	func forwardTraffic(clientConn net.Conn,config *SSConfig){
		//处理数据并转发
		defer clientConn.Close()
		remoteConn,err := connectToRemote(config)
		defer remoteConn.Close()
		if err != nil{
			log.Fatalf("发送数据失败")
		}
		log.Println(" 代理建立TCP成功连接%d",config.Server)  
		buf := make([]byte,4096) 
		for {
			
			n,err := clientConn.Read(buf)
			//Read() 方法会尝试从 conn 读取数据，并将其存储到 （buf）中
			if err !=nil {
				//表示数据读取完毕了
				break
			}
			log.Println(" 从客户端读取数据")
			encryptrData,_ := EnctyptAES256GCM(buf[:n],[]byte(config.Password))
			//进行加密 
			remoteConn.Write(encryptrData)
			log.Println(" 数据成功发送",config.ServerPort)
		}
	}

	