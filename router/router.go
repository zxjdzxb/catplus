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
	//tags 增删改查
	r.POST("/api/v1/tags", middleware.AuthMiddleware(), controllor.CreateTagHandler)
	r.PATCH("/api/v1/tags/:id", controllor.UpdateTagHandler)
	r.GET("/api/v1/tags/:id", controllor.GetTagHandler)
	r.DELETE("/api/v1/tags/:id", controllor.DeleteTagHandler)
	r.GET("/api/v1/tags", controllor.GetTagListHandler)
	//items 增删改查
	r.POST("/api/v1/items", controllor.CreateItemHandler)
	r.GET("/api/v1/items/summary", controllor.GetItemsSummaryHandler)
	r.GET("/api/v1/items", controllor.GetItemsHandler)

	return r
}
