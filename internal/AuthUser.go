package internal

import (
	"fmt"
	"net/smtp"
)

func Login() {
	// 邮箱服务器的地址和端口
	host := "smtp.qq.com"
	port := 587

	// 邮箱账号和密码
	email := "1660154581@qq.com"
	password := "ngtyjjpxzpbkbeea"

	// 使用 PlainAuth 进行邮箱登录
	auth := smtp.PlainAuth("", email, password, host)

	// 连接邮箱服务器
	client, err := smtp.Dial(fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println("连接邮箱服务器失败：", err)
		return
	}

	// 发送 EHLO 命令
	if err := client.Hello("localhost"); err != nil {
		fmt.Println("发送 EHLO 命令失败：", err)
		return
	}

	// 使用 TLS 加密连接
	if err := client.StartTLS(nil); err != nil {
		fmt.Println("使用 TLS 加密连接失败：", err)
		return
	}

	// 验证邮箱账号和密码
	if err := client.Auth(auth); err != nil {
		fmt.Println("验证邮箱账号和密码失败：", err)
		return
	}

	fmt.Println("登录成功！")
}
