package utils

import "unicode"

// ValidatePassword
// 1. 长度 >= 9位
// 2. 至少包含一个大写字母
// 3. 至少包含一个小写字母
// 4. 至少包含一个数字
func ValidatePassword(password string) bool {
	// 首先检查长度
	if len(password) < 9 {
		return false
	}

	// 初始化验证标志
	var hasUpper, hasLower, hasDigit bool

	// 遍历每个字符进行校验
	for _, c := range password {
		switch {
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		case unicode.IsDigit(c):
			hasDigit = true
		}

		// 提前退出优化：当所有条件都满足时立即返回
		if hasUpper && hasLower && hasDigit {
			break
		}
	}

	return hasUpper && hasLower && hasDigit
}
