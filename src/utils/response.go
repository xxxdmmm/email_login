package utils

import "github.com/gin-gonic/gin"

func UserBad(context *gin.Context, msg string, data interface{}) {
	context.JSON(400, gin.H{
		"status":  false,
		"message": msg,
		"data":    data,
	})
}

func Success(context *gin.Context, msg string, data interface{}) {
	context.JSON(200, gin.H{
		"status":  true,
		"message": msg,
		"data":    data,
	})
}

func ServerBad(context *gin.Context, msg string, data interface{}) {
	context.JSON(500, gin.H{
		"status":  false,
		"message": msg,
		"data":    data,
	})
}
