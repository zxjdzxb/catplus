package main

import (
	"catplus-server/database"
	"catplus-server/middleware"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Email              string `gorm:"unique"`
	Password           string
	VerificationCode   string
	VerificationExpiry time.Time
}

func main() {
	InitConfig()

	// 连接到 MySQL 数据库
	db := database.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 自动迁移 User 模型
	db.AutoMigrate(&User{})

	// 创建 Gin 引擎
	r := gin.Default()

	// 处理生成验证码的请求
	r.POST("/generate-code", func(c *gin.Context) {
		// 从请求中获取电子邮件地址
		email := c.PostForm("email")

		// 从数据库中查找用户
		var user User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}

		// 生成 6 位随机验证码
		code := make([]byte, 3)
		if _, err := rand.Read(code); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate code"})
			return
		}
		codeStr := hex.EncodeToString(code)

		// 将验证码保存到数据库
		expiry := time.Now().Add(time.Minute * 10) // 验证码有效期为 10 分钟
		if err := db.Model(&user).Updates(User{
			VerificationCode:   codeStr,
			VerificationExpiry: expiry,
		}).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save code"})
			return
		}

		// 发送电子邮件
		middleware.Message()

		c.Status(http.StatusOK)
	})

	// 处理验证验证码的请求
	r.POST("/verify-code", func(c *gin.Context) {
		// 从请求中获取电子邮件地址和验证码
		email := c.PostForm("email")
		code := c.PostForm("code")

		// 从数据库中查找用户
		var user User
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

		c.Status(http.StatusOK)
	})

	// 运行 Gin 应用程序
	if err := r.Run(":8080"); err != nil {
		panic("failed to start server")
	}
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
