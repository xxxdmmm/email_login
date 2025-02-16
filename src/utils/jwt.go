package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var mySigningKey = []byte("this_is_your_secret")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(username string) (string, error) {
	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "xxxdmmm",
			Subject:   "for_auth",
		},
	}

	// 选择签名算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成token字符串
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

// 验证 Token 函数
func validateToken(tokenStr string) (*MyCustomClaims, error) {

	// 解析 Token 并验证，回调函数返回的是JWT密钥与错误信息
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	// 如果解析失败，返回错误
	if err != nil {
		return nil, err
	}

	// 检查 token 是否有效
	if !token.Valid {
		// 如果 Token 无效，返回错误
		return nil, errors.New("invalid or expired token")
	}

	// 尝试将 token.Claims 转换为 *Claims 类型
	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok {
		// 如果类型断言失败，返回错误
		return nil, errors.New("invalid claims")
	}

	// 如果以上检查都通过，返回解析后的 claims
	return claims, nil
}
