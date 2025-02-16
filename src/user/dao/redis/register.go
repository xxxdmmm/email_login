package redis

import (
	"awesomeProject/src/utils"
	"context"
	"log"
	"time"
)

const (
	FOR_REGISTER = "user:register:email:"
	FOR_LOGIN    = "user:login:email:"
)

func SaveEmailCodeToRedis(email string, code string, t string) error {
	redisEngine := utils.DefaultRedis
	ct := context.Background()

	// 3分钟有效
	res := redisEngine.Set(ct, t+email, code, time.Minute*3)

	if res.Err() != nil {
		log.Default().Println("Failed to save email to Redis:", res.Err())
		return res.Err()
	}

	return nil
}

func CheckEmailInRedis(email string, t string) bool {
	redisEngine := utils.DefaultRedis
	ct := context.Background()
	exists, err := redisEngine.Exists(ct, t+email).Result()

	if err != nil {
		log.Default().Println("Failed to check email in Redis:", err)
		return false
	}

	// 如果存在，则返回true
	if exists == 1 {
		return true
	}
	return false
}

func CheckCode(email string, t string) string {
	redisEngine := utils.DefaultRedis
	ct := context.Background()

	res := redisEngine.Get(ct, t+email)

	value, err := res.Result()

	if err != nil {
		log.Default().Println("Failed to get email from Redis:", err)
		return ""
	}

	return value
}
