package main

import (
	"awesomeProject/src/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// 创建一个服务
	app := gin.Default()

	app.Use(cors.Default())

	user.User(app)

	// 监听8080端口
	err := app.Run(":8080")

	if err != nil {
		log.Fatal("Error:", err)
	}
}
