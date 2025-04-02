package main 
import (
   SS5  "Shadowsock/SS5"
)
func main(){
	config :=SS5.DefaultConfig() 
	SS5.StartSocks5(config)
}