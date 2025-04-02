package SS5
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EnctyptAES256GCM(plaintext []byte ,key []byte )([]byte,error){ //plaintext需要加密的明文数据 ，key 密钥 32字节
	block,err := aes.NewCipher(key)    //按照密钥生成分组密钥长度必须满足32字节
	if err != nil {
		return nil ,err 
	}
	//Block 里面有加密和解密的方法 它是一个接口类型  它是属于cipher.Block——生成密钥
	aesGCM , err := cipher.NewGCM(block) 
	//ase转化为GCM 模式并且使他支持认证加密——保证数据不会被明文泄露
	nonce := make([]byte, aesGCM.NonceSize())
	//生成随机数确保数据加密的唯一的随机数，防止重放攻击，不需要加密的
	if _,err := io.ReadFull(rand.Reader , nonce );err != nil{
		return nil ,err 
	} 
	//生成安全的随机数填充nonce 
	ciphertext := aesGCM.Seal(nil,nonce ,plaintext,nil)
	//只是对内容进行了加密并没有附带nonce
	result := append(nonce, ciphertext...)
	//将文件进行加密并且附带nonce防止重放攻击
	return result ,nil 
}

