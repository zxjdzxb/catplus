package controllor

import (
	"catplus-server/database"
	"catplus-server/middleware"
	"catplus-server/model"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// VerificationCode 生成验证码

func VerificationCode(c *gin.Context) {
	email := c.PostForm("email")
	db := database.GetDB()

	// 从数据库中查找用户,若没有则创建
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		if err := db.Create(&model.User{Email: email}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}
	}

	// // 生成 6 位随机验证码
	codeStr, err := middleware.GenerateVerificationCode()
	if err != nil {
		log.Println("生成验证码失败：", err)
		return
	}

	//将验证码保存到数据库

	expiry := time.Now().Add(time.Minute * 10) // 验证码有效期为 10 分钟
	log.Println("expiry:", time.Now())
	// newUser := model.User{
	// 	Email:              email,
	// 	VerificationCode:   codeStr,
	// 	VerificationExpiry: expiry,
	// }

	if err := db.Model(&user).Updates(model.User{
		VerificationCode:   codeStr,
		VerificationExpiry: expiry,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save code"})
		return
	}

	// 发送电子邮件
	if !middleware.IsValidEmail(email) {
		log.Println("邮箱格式不正确")
		return
	}
	if err := middleware.SendVerificationCodeToEmail(email, codeStr); err != nil {
		log.Println("发送邮件失败：", err)
		return
	}
	log.Println("验证码已发送到邮箱，请注意查收")

	c.JSON(http.StatusOK, gin.H{"message": "verification code sent"})

}

// VerifyCode 验证验证码
func VerifyCode(c *gin.Context) {
	email := c.PostForm("email")
	code := c.PostForm("code")
	db := database.GetDB()
	// 从数据库中查找用户
	var user model.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// 检查验证码是否过期
	if user.VerificationExpiry.Before(time.Now()) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "verification code has expired"})
		return
	}

	// 检查验证码是否正确
	if user.VerificationCode != code {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid verification code"})
		return
	}

	// 验证成功，可以执行其他操作
	// ...

	c.JSON(http.StatusOK, gin.H{"message": "verification successful"})
}
