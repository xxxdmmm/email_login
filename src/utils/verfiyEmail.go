package utils

import (
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`)

func ValidateEmail(email string) (bool, string) {
	// 去除前后空格
	email = strings.TrimSpace(email)

	// 基本长度检查
	if len(email) < 6 || len(email) > 254 {
		return false, email
	}

	// 正则表达式验证
	return emailRegex.MatchString(email), email
}
