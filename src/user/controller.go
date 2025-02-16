package user

import (
	"awesomeProject/src/user/server"
	"github.com/gin-gonic/gin"
)

func User(app *gin.Engine) {
	/*
		获取验证码接口
		提供的信息：
			邮箱
		请求示例：
			{
				email: "123@qq.com"
			}
	*/
	app.POST("/email/code/:type", server.GetEmailCode)

	/*
		注册接口
		提供的信息：
			邮箱
			密码
			验证码
		请求示例：
			{
				email: "123@qq.com",
				password: "1312",
				code: "234567"
			}
	*/
	app.POST("/user/register", server.Register)

	/*
		登录接口
		提供信息：
			邮箱
			验证码

	*/
	app.POST("/user/login/code", server.LoginByCode)

	/*
		登录接口
		提供信息：
			邮箱
			密码
	*/
	app.POST("/user/login/password", server.LoginByPassword)
}
