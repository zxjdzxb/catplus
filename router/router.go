package router

import (
	"catplus-server/controllor"
	"catplus-server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/v1/validation_codes", controllor.VerificationCode)
	r.POST("/api/v1/session", controllor.VerifyCode)
	r.GET("/api/v1/me", middleware.AuthMiddleware(), controllor.Info)
	r.POST("/api/v1/tags", controllor.CreateTag)

	return r
}
