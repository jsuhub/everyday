package main
import (
	"log"
	"net"
	"fmt"
	"encoding/binary"
)


var conn_socks5_server *net.Conn 

func main(){
listener,err := net.Listen("tcp","0.0.0.0:1080")
defer listener.Close()
if err != nil {
log.Fatalf("启动监听失败:%v",err)
}
log.Println("监听端口 1080...")
for {
conn,err := listener.Accept()
if err != nil {
log.Printf("接受连接失败:v",err)
continue
}
go handleConnection(conn)
}
}

func handleConnection(conn net.Conn){
	defer conn.Close()
// 这个conn表示的是一个请求发送过来了
for{
	buf := make([]byte,4096)
	n,err := conn.Read(buf)
	//这里就相当于进行TCP连接不能中断
		if err != nil {
		log.Fatalf("读取TCP连接内容失败")}
		buf = buf[:n]
		if buf[0] != 0x05 {
		log.Printf("不是 S0CKS5 协议，收到的版本号:%d", buf[0])
		return
		}
		var VER byte =0x05
		buf1 := make([]byte,2)
		if n< 8 {
			log.Println("第一次请求")
		var METHOD  byte = 0x00
		buf1[0] = VER 
		buf1[1] = METHOD
		}else{
			buf1 = make([]byte,4)
			var RSV byte = 0x00
			var REP byte= 0x00 
			ATYP := buf[3]
			PORT := make([]byte,2)
			PORT = buf[n-2:n]

			log.Printf("%x",PORT)
			
			BND := make([]byte,1024) 
			var serverAddr string 
			switch ATYP {
			case 0x01:
				BND = buf[4:9]
				BND = BND[:5]
				a := int(BND[0])
				b := int(BND[1])
				c := int(BND[2])
				d := int(BND[3])
				serverAddr = fmt.Sprintf("%d.%d.%d.%d",a,b,c,d)
			//代表ipv4的报文
			case 0x04 : 
				BND = buf[4:20]
				BND = BND[:16]
				log.Println(BND)
		

			default : 
				BND = buf[4:n-2]
				BND = BND [:n-2-4]
				serverAddr = string(BND)
			}
			num := binary.BigEndian.Uint16(PORT)
			serverAddr = fmt.Sprintf("%s:%d",serverAddr,num)
						log.Printf(serverAddr)
			conn2 ,err := net.Dial("tcp",serverAddr)
			if err != nil{
				log.Fatalf("Socks5与目的服务器建立TCP连接失败%v",err)
			}
			conn_socks5_server = &conn2
			log.Println(serverAddr)
			buf1[0] = VER 
			buf1[1] = REP
			buf1[2] = RSV
			buf1[3] = ATYP 
			buf1 = append(append(buf1,BND...),PORT...)
	}
		conn.Write(buf1)
		log.Println(buf1)
}
}
