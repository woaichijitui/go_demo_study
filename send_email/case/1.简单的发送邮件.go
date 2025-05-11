package _case

import (
	"encoding/hex"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
)

// SMTP configuration
const (
	SmtpHost = "smtp.qq.com"
	SmtpPort = 587
	SmtpUser = "1975611740@qq.com"
	SmtpPass = "xjerlkpwgahacjhf"
)

// EmailTokenMap 用于存储邮箱和令牌的映射
var EmailTokenMap = make(map[string]string)

// GenerateToken 生成唯一令牌
func GenerateToken() (string, error) {
	token := make([]byte, 16)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// SendEmail 发送邮件
func SendEmail(email, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", SmtpUser)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "htt-wep 登录验证") //邮件主题
	m.SetBody("text/plain", fmt.Sprintf("登录验证码：%s", token))

	d := gomail.NewDialer(SmtpHost, SmtpPort, SmtpUser, SmtpPass)

	return d.DialAndSend(m)
}
