package mysql

import (
	"awesomeProject/src/user/entry"
	"awesomeProject/src/utils"
	"log"
)

func CheckRepeat(email string) bool {
	mysql := utils.DefaultEngine

	var user entry.Users

	has, err := mysql.Engine.Where("email = ?", email).Get(&user)

	if err != nil {
		log.Default().Println("Failed to check email in MySQL:", err)
	}

	// 已存在则返回true
	return has
}

func SaveNewUser(user *entry.Users) int64 {
	mysql := utils.DefaultEngine

	num, err := mysql.Engine.Insert(user)

	if err != nil || num == 0 {
		log.Default().Println("Failed to save new user in MySQL:", err)
	}

	return num
}
