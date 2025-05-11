package testdata

import (
	_case "send_email/case"
	"testing"
)

func TestSendEmail(t *testing.T) {
	token, _ := _case.GenerateToken()
	_case.SendEmail("1975611740@qq.com", token)
}
