package router

import (
	"catplus-server/controllor"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/generate-code", controllor.VerificationCode)
	r.POST("/verify-code", controllor.VerifyCode)

	return r
}
