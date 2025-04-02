 package SS5

// import (
// 	"crypto/md5"
// 	"encoding/hex"
// )

// // EVP_BytesToKey 生成密钥和 IV
// func EVP_BytesToKey(password string, keyLen int, ivLen int, salt []byte) ([]byte, []byte) {
// 	var key, iv []byte
// 	var prev []byte

// 	// 计算 key 和 iv
// 	for len(key) < keyLen || len(iv) < ivLen {
// 		md5Hash := md5.New()
// 		md5Hash.Write(prev)     // 上次的 MD5 结果
// 		md5Hash.Write([]byte(password)) // 用户密码
// 		if salt != nil {
// 			md5Hash.Write(salt) // 可选的 salt
// 		}
// 		prev = md5Hash.Sum(nil)

// 		if len(key) < keyLen {
// 			key = append(key, prev...)
// 		} else {
// 			iv = append(iv, prev...)
// 		}
// 	}

// 	return key[:keyLen], iv[:ivLen]
// }

// func productKey(password []byte) ([]byte,[]byte) {
// 	salt := []byte("shadowsocks") // 这里可以替换成随机 salt
// 	keyLen := 32
// 	ivLen := 16

// 	return key, iv := EVP_BytesToKey(password, keyLen, ivLen, salt)
	 
// }
