package middleware

import (
	"bytes"
	"catplus-server/common"
	"catplus-server/database"
	"catplus-server/model"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// 发送验证码到指定邮箱
func SendVerificationCodeToEmail(email string, code string) error {

	// 邮件正文内容
	fromName := "zxjdzxb"
	subject := "CatPlus 验证码"

	// SMTP 服务器地址
	smtpServer := "smtp.qq.com"

	// 发送者邮箱账号和密码
	from := viper.GetString("email.username")
	password := viper.GetString("email.password")

	// 收件人邮箱
	to := email
	//
	log.Println("to:", to)
	// message := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", fromName, from, to, subject, code)
	message := generateEmailMessage(fromName, from, to, subject, code)
	// SMTP 认证信息
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// 发送邮件
	err := smtp.SendMail(smtpServer+":587", auth, from, []string{to}, []byte(message))
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

func generateEmailMessage(fromName, from, to, subject, code string) string {
	// 定义邮件模板
	const templateText = `<html>
	<head>
		<meta charset="utf-8">
		<title>{{.Subject}}</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f2f2f2;
				padding: 20px;
				margin: 0;
			}

			.container {
				max-width: 600px;
				margin: 0 auto;
				background-color: #fff;
				border: 1px solid #ddd;
				padding: 20px;
			}

			h1 {
				font-size: 24px;
				color: #333;
				margin-bottom: 20px;
			}

			p {
				font-size: 16px;
				line-height: 1.5;
				color: #555;
			}

			strong {
				font-weight: bold;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>{{.Subject}}</h1>
			<p>Hello!</p>
			<p>验证码: <strong>{{.Code}}</strong></p>
			<p>有效期十分钟.</p>
		</div>
	</body>
	</html>`

	// 解析邮件模板
	tmpl, err := template.New("email").Parse(templateText)
	if err != nil {
		panic(err)
	}

	// 渲染邮件模板
	data := struct {
		FromName string
		From     string
		To       string
		Subject  string
		Code     string
	}{
		FromName: fromName,
		From:     from,
		To:       to,
		Subject:  subject,
		Code:     code,
	}
	var messageBody bytes.Buffer
	err = tmpl.Execute(&messageBody, data)
	if err != nil {
		panic(err)
	}

	// 生成邮件消息
	message := fmt.Sprintf("From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"utf-8\"\r\n\r\n%s", fromName, from, to, subject, messageBody.String())
	return message
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "token签名无效",
				})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"code": 401,
					"msg":  "无效的token",
				})
			}
			c.Abort()
			return
		}

		// 验证token是否过期
		if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "token已过期",
			})
			c.Abort()
			return
		}

		userId := claims.UserId
		DB := database.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "用户不存在",
			})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
