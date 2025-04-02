package main 
import (
	SS5  "Shadowsock_Server/SS5"
 )
 func main(){
	 config :=SS5.DefaultConfig() 
	 SS5.StartSocks5_server(config)
 }