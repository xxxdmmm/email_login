package server

import (
	"awesomeProject/src/user/dao/mysql"
	"awesomeProject/src/user/dao/redis"
	"awesomeProject/src/utils"
	"github.com/gin-gonic/gin"
	"time"
)

type UserWithPassword struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserWithCode struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func LoginByCode(context *gin.Context) {
	var user UserWithCode
	if err := context.ShouldBindJSON(&user); err != nil {
		utils.UserBad(context, "请完善字段信息", nil)
		return
	}

	// 校验邮箱格式
	verify, email := utils.ValidateEmail(user.Email)
	if !verify {
		utils.UserBad(context, "邮箱格式错误", nil)
		return
	}

	// 判断验证码是否存在于redis
	if !redis.CheckEmailInRedis(email, redis.FOR_LOGIN) {
		utils.UserBad(context, "未发送验证码", nil)
		return
	}

	// 校验验证码
	if redis.CheckCode(email, redis.FOR_LOGIN) != user.Code {
		utils.UserBad(context, "验证码错误", nil)
		return
	}

	// 检查是否存在
	check, u := mysql.LoginByCode(email)

	if !check {
		utils.UserBad(context, "用户不存在", nil)
		return
	}

	token, err := utils.GenerateJWT(u.Email)

	if err != nil {
		utils.ServerBad(context, "服务器错误", nil)
		return
	}

	utils.Success(context, "登录成功", gin.H{
		"id":    u.Id,
		"name":  u.Name,
		"email": u.Email,
		"token": token,
		"time":  time.DateTime,
	})
}

func LoginByPassword(context *gin.Context) {
	var user UserWithPassword
	if err := context.ShouldBindJSON(&user); err != nil {
		utils.UserBad(context, "请完善字段信息", nil)
		return
	}

	// 验证邮箱格式
	check, email := utils.ValidateEmail(user.Email)
	if !check {
		utils.UserBad(context, "邮箱格式错误", nil)
		return
	}

	// 验证密码格式
	check = utils.ValidatePassword(user.Password)
	if !check {
		utils.UserBad(context, "用户名或密码错误", nil)
		return
	}

	// 验证密码是否正确
	has, u := mysql.LoginByPassword(email, user.Password)

	if !has {
		utils.UserBad(context, "用户名或密码错误", nil)
		return
	}

	token, err := utils.GenerateJWT(u.Email)

	if err != nil {
		utils.ServerBad(context, "服务器错误", nil)
		return
	}

	utils.Success(context, "登录成功", gin.H{
		"id":    u.Id,
		"name":  u.Name,
		"email": u.Email,
		"token": token,
		"time":  time.DateTime,
	})
}
