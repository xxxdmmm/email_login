package utils

import (
	"gopkg.in/gomail.v2"
	"log"
)

func SendEmail(email string, code string) {
	sender := ""
	password := ""
	smtpHost := "Smtp.qq.com"
	smtpPort := 587

	m := gomail.NewMessage()

	m.SetHeader("From", sender)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/html", "您的验证码为：<b>"+code+"</b>!")

	// 发送邮件
	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)

	err := d.DialAndSend(m)

	if err != nil {
		log.Default().Println("Failed to send email:", err)
		return
	}
	log.Default().Println("Email sent successfully!")
}
