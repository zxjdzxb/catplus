package controllor

import (
	"catplus-server/common"
	"catplus-server/database"
	"catplus-server/middleware"
	"catplus-server/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

/*
Content-Type: application/json
{
	"email": 1660154581@qq.com
}
*/
// VerificationCode 生成验证码
func VerificationCode(c *gin.Context) {
	type RequestBody struct {
		Email string `json:"email"`
	}

	var reqBody RequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		common.Fail(c, gin.H{}, "invalid request body")
		return
	}

	db := database.GetDB()

	// 从数据库中查找用户,若没有则创建
	var user model.User
	if err := db.Where("email = ?", reqBody.Email).First(&user).Error; err != nil {
		if err := db.Create(&model.User{Email: reqBody.Email}).Error; err != nil {
			common.Fail(c, gin.H{}, "failed to create user")
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
	if err := db.Model(&user).Updates(model.User{
		VerificationCode:   codeStr,
		VerificationExpiry: expiry,
	}).Error; err != nil {
		common.Fail(c, gin.H{}, "failed to save code")
		return
	}

	// 发送电子邮件
	if !middleware.IsValidEmail(reqBody.Email) {

		common.Fail(c, gin.H{}, "无效的邮箱地址")
		return
	}
	if err := middleware.SendVerificationCodeToEmail(reqBody.Email, codeStr); err != nil {
		log.Println(err)
		common.Fail(c, gin.H{}, "发送邮件失败")
		return
	}

	common.Success(c, gin.H{}, "验证码已发送到邮箱，请注意查收")

}

// VerifyCode 验证验证码
func VerifyCode(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		common.Fail(c, gin.H{"error": err.Error()}, "invalid request body")
		return
	}
	db := database.GetDB()
	// 从数据库中查找用户
	var user model.User
	if err := db.Where("email = ?", request.Email).First(&user).Error; err != nil {
		common.Fail(c, gin.H{}, "user not found")
		return
	}

	// 检查验证码是否过期
	if user.VerificationExpiry.Before(time.Now()) {
		common.Fail(c, gin.H{}, "verification code has expired")
		return
	}

	// 检查验证码是否正确
	if user.VerificationCode != request.Code {
		common.Fail(c, gin.H{}, "invalid verification code")
		return
	}

	// 验证成功，可以执行其他操作
	token, err := common.ReleaseToken(user)
	if err != nil {
		common.Fail(c, gin.H{
			"code": 500,
		}, "系统异常")
		log.Printf("token generate error : %v", err)
		return
	}

	common.Success(c, gin.H{"token": token}, "验证成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	common.Success(c, gin.H{"user": user}, "获取用户信息成功")
}
