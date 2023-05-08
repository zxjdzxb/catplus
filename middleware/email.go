package middleware

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"regexp"
	"strconv"
)

// 发送验证码到指定邮箱
func SendVerificationCodeToEmail(email string) error {
	// 生成随机的验证码
	code, err := GenerateVerificationCode()
	if err != nil {
		log.Println("生成验证码失败：", err)
		return err
	}
	// 邮件正文内容
	fromName := "Sender Name"
	subject := "Test Email"

	// SMTP 服务器地址
	smtpServer := "smtp.qq.com"

	// 发送者邮箱账号和密码
	from := "1660154581@qq.com"
	password := "ngtyjjpxzpbkbeea"

	// 收件人邮箱
	to := []string{email}
	log.Println("to:", to)
	message := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", fromName, from, to, subject, code)

	// SMTP 认证信息
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// 发送邮件
	err = smtp.SendMail(smtpServer+":587", auth, from, to, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

// generateVerificationCode 生成 6 位随机数字
func GenerateVerificationCode() (string, error) {

	code := rand.Intn(900000) + 100000
	return strconv.Itoa(code), nil
}

// 验证邮箱格式是否正确
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func Message(emai string) {
	email := emai
	if !IsValidEmail(email) {
		fmt.Println("邮箱格式不正确")
		return
	}
	if err := SendVerificationCodeToEmail(email); err != nil {
		fmt.Println("发送邮件失败：", err)
		return
	}
	fmt.Println("验证码已发送到邮箱，请注意查收")
}
