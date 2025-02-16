package server

import (
	"awesomeProject/src/user/dao/mysql"
	"awesomeProject/src/user/dao/redis"
	"awesomeProject/src/user/entry"
	"awesomeProject/src/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

type User struct {
	Email    string `json:"email" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(context *gin.Context) {
	var user User

	err := context.Bind(&user)

	if err != nil {
		utils.UserBad(context, "请完善信息", nil)
		return
	}

	// 校验密码格式
	if !utils.ValidatePassword(user.Password) {
		utils.UserBad(context, "密码格式错误", nil)
		return
	}

	// 校验邮箱格式
	verify, email := utils.ValidateEmail(user.Email)
	if !verify {
		utils.UserBad(context, "邮箱格式错误", nil)
		return
	}

	// 判断验证码是否存在于redis
	if !redis.CheckEmailInRedis(email, redis.FOR_REGISTER) {
		utils.UserBad(context, "未发送验证码", nil)
		return
	}

	// 校验验证码
	if redis.CheckCode(email, redis.FOR_REGISTER) != user.Code {
		utils.UserBad(context, "验证码错误", nil)
		return
	}

	// 检查是否存在
	if mysql.CheckRepeat(email) {
		utils.UserBad(context, "用户已存在", nil)
		return
	}

	// 保存
	num := mysql.SaveNewUser(&entry.Users{
		Email:    email,
		Password: user.Password,
	})

	if num == 0 {
		utils.ServerBad(context, "注册失败", nil)
		return
	}

	utils.Success(context, "注册成功", nil)
}

func GetEmailCode(context *gin.Context) {
	typeOfSend := context.Param("type")

	if typeOfSend == "register" {
		typeOfSend = redis.FOR_REGISTER
	} else if typeOfSend == "login" {
		typeOfSend = redis.FOR_LOGIN
	} else {
		utils.UserBad(context, "请求错误", nil)
		return
	}

	data, err := context.GetRawData()

	if err != nil {
		utils.UserBad(context, "请求错误", nil)
		return
	}

	var JSONData map[string]interface{}

	// 反序列化
	err = json.Unmarshal(data, &JSONData)

	if err != nil {
		utils.UserBad(context, "请求错误", nil)
		return
	}

	email, ok := JSONData["email"]

	if !ok {
		utils.UserBad(context, "请填写邮箱", nil)
		return
	}

	emailVal, ok := email.(string)

	if !ok {
		utils.ServerBad(context, "请求错误，请联系管理员1", nil)
		return
	}

	// 验证邮箱格式
	verify, emailVal := utils.ValidateEmail(emailVal)
	if !verify {
		utils.UserBad(context, "邮箱格式错误", nil)
		return
	}

	// 检查是否三分钟内重复请求
	if redis.CheckEmailInRedis(emailVal, typeOfSend) {
		utils.UserBad(context, "请勿频繁发送验证码", nil)
		return
	}

	// 生成随机的6位数字
	randomNumber := rand.Intn(900000) + 100000
	number := fmt.Sprintf("%06d", randomNumber)

	// 保存到Redis
	if err := redis.SaveEmailCodeToRedis(emailVal, number, typeOfSend); err != nil {
		utils.ServerBad(context, "请求错误，请联系管理员2", nil)
		return
	}

	// 发送邮件
	go utils.SendEmail(emailVal, number)

	utils.Success(context, "验证码发送成功，三分钟内有效", nil)
}
