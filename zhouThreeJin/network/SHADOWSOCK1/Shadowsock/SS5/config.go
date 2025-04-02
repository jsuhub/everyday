package SS5

type SSConfig struct {
	Server string 
	ServerPort int 
	LocalPort int 
	Password string 
	Method string 
}


func DefaultConfig () *SSConfig{
	return &SSConfig{
		Server :	"127.0.0.1",  //服务器地址  
		ServerPort : 18888  ,
		LocalPort : 18889  ,//监听本地端口的数据——也就是客户端流出的那种流量需要被走代理了  
		Password : "12345678901234567890123456789012",
		Method : "ase-256-gcm",
	}
}

