package mysql

import (
	"awesomeProject/src/user/entry"
	"awesomeProject/src/utils"
	"log"
)

func LoginByPassword(email, password string) (bool, *entry.Users) {
	mysql := utils.DefaultEngine
	var user entry.Users
	has, err := mysql.Engine.Where("email=?", email).Where("password=?", password).Get(&user)

	if err != nil {
		log.Default().Println("Failed to check email in MySQL:", err)
	}

	return has, &user
}

func LoginByCode(email string) (bool, *entry.Users) {
	mysql := utils.DefaultEngine
	var user entry.Users
	has, err := mysql.Engine.Where("email=?", email).Get(&user)

	if err != nil {
		log.Default().Println("Failed to check email in MySQL:", err)
	}

	return has, &user
}
