package main  
import(
	"fmt"
	"net"   //提供了许多网络I/O接口的方法
	"io"
)
func process(conn net.Conn){  //这个是协程里面处理数据的
defer conn.Close() //用来关闭连接/管道，需要释放连接。不然会卡住    
for {
	//每次都要创建新的切片
	buf := make([]byte,1024) 
	//如果通道里面没有Write东西就会一直等待,或者超时  
	n,err := conn.Read(buf)    //用来读取Write的数据并写入buf中
	if err != nil{
		fmt.Println("服务器的Read err = ",err)
		return //表示一旦对面服务器关闭了，那么这个for循环也应该退出
	}
	//显示客户端的数据在服务器终端  
	//客户端输入的时候\n已经被传进来了，所以最后有一个换行
	fmt.Print(string(buf[:n]))//这个n代表真正东西的长度
}

}

func main(){
	fmt.Println("服务器开始监听")  
	//tcp 表示协议,这里表示只监听一下，listen是一个接口
	listen,err := net.Listen("tcp","0.0.0.0:8888") //直接调用net包里面的内容,如果是ipv6可以用0.0.0.0表示本机,注意如果其他来监听最好使用0.0.0.0如何使用127.0.0.1可能会出现监听失败
	if err !=nil{
		//	表示监听失败没有数据传输过来  
		fmt.Println("listen_err :",err)	
		return 
	}	
	defer listen.Close()  //表示最后关闭这个接口，defer关键字的用处  
     //循环等待客户机传来链接  
	 for {
		fmt.Println("wait client chain")
		conn,err :=listen.Accept() //这个函数是用来等待客户端链接
	    if err== io.EOF {
			fmt.Println("Accept() = ",err)
			//存在错误是可以的，不需要中断循环
			} else {
			//表示一个连接成功  telnet可以用来测试端口  如：telnet www.baidu.com  80  输入ctrl+]退出
			fmt.Println("Accept() = %v\n客户端ip=%v",conn,conn.RemoteAddr().String())
			 process(conn)
		
		}
	}
	//客户端发来请求现在应该想办法生成一个新的线程不能让他阻塞  


}
