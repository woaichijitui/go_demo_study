package _case

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// 解析jwt字符串
func ParseJwt(key any, jwtStr string, options ...jwt.ParserOption) (claim jwt.Claims, err error) {
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, options...)
	if err != nil {
		return nil, err
	}

	// 校验 Claims 对象是否有效，基于 exp（过期时间），nbf（不早于），iat（签发时间）等进行判断（如果有这些声明的话）
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return token.Claims, nil
}
