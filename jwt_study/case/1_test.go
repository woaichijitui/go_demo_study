package _case

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"testing"
)

func TestGenerateJwt(T *testing.T) {
	jwtPwd := make([]byte, 32) // 生成32字节（256位）的密钥
	if _, err := rand.Read(jwtPwd); err != nil {
		panic(err)
	}

	/*iss(Issuer)	发行者，标识 JWT 的发行者。
	sub(Subject)	主题，标识 JWT 的主题，通常指用户的唯一标识
	aud(Audience)	观众，标识 JWT 的接收者
	exp(Expiration Time)	过期时间。标识 JWT 的过期时间，这个时间必须是将来的
	nbf(Not Before)	不可用时间。在此时间之前，JWT 不应被接受处理
	iat(Issued At)	发行时间，标识 JWT 的发行时间
	jti(JWT ID)	JWT 的唯一标识符，用于防止 JWT 被重放（即重复使用）

	*/
	jwtStr, err := GenerateJwt(jwtPwd, jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "htt",
		"sub": "aichijitui.cn",
		"aud": "Programmer",
		"exp": 4,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(jwtStr)

}
