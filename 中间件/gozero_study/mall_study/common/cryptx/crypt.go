package cryptx

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
)

// PasswordEncrypt 使用 scrypt 算法对密码进行加密。
// scrypt 是一种内存硬函数，适用于密码存储，能有效抵御暴力破解攻击。
// 参数 salt 是一个随机字符串，用于增加加密的安全性，防止彩虹表攻击。
// 参数 password 是用户输入的原始密码。
// 返回值为加密后的十六进制字符串。
func PasswordEncrypt(salt, password string) string {
	// 使用 scrypt.Key 函数对密码进行加密。
	// 32768 是 N 参数，表示 CPU/内存成本参数，值越大安全性越高但性能越低。
	// 8 是 r 参数，表示块大小参数。
	// 1 是 p 参数，表示并行化参数。
	// 32 是期望的密钥长度。
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	// 将加密后的字节切片转换为十六进制字符串并返回。
	return fmt.Sprintf("%x", string(dk))
}
