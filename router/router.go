package router

import (
	"catplus-server/controllor"
	"catplus-server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/generate-code", controllor.VerificationCode)
	r.POST("/verify-code", controllor.VerifyCode)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controllor.Info)

	return r
}
