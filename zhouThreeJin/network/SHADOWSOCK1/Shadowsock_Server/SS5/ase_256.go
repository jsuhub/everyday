package SS5
import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func DecryptAES256GCM(ciphertext []byte,key []byte )([]byte,error){
	block,err :=aes.NewCipher(key) //密钥分段
	if err != nil {
		return nil ,err 
	}
	aesGCM,err := cipher.NewGCM(block) //密钥转化为GCM模式
	if err != nil {
		return nil ,err 
	}
	nonceSize := aesGCM.NonceSize()			//使用该加密方式生成随机数以及填充的大小		
	if len(ciphertext)<nonceSize {
		return nil ,errors.New("填充数字大于加密总体大小") 
	}  
	nonce , ciphertdata := ciphertext[:nonceSize],ciphertext[nonceSize:]  //将内容和填充分开，填充的在前面 
	plaintext,err := aesGCM.Open(nil,nonce,ciphertdata,nil) //利用aes和密钥解密  
	if err != nil{
		return nil,err 
	}
	return plaintext,nil 

}