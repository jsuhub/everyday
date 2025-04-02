package SS5

 type SSConfig struct {
	LocalPort int 
	Password string 
	Method string 
}

func DefaultConfig() *SSConfig{
	return &SSConfig{
		LocalPort : 18888 ,
		Password : "12345678901234567890123456789012",
		Method : "ase-256-gcm",
	}
}