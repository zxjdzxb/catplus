package router

import (
	"catplus-server/controllor"
	"catplus-server/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/validation_codes", controllor.VerificationCode)
		v1.POST("/session", controllor.VerifyCode)
		v1.GET("/me", middleware.AuthMiddleware(), controllor.Info)

		tags := v1.Group("/tags") //, middleware.AuthMiddleware()
		{
			tags.POST("", controllor.CreateTagHandler)
			tags.PATCH("/:id", controllor.UpdateTagHandler)
			tags.GET("/:id", controllor.GetTagHandler)
			tags.DELETE("/:id", controllor.DeleteTagHandler)
			tags.GET("", controllor.GetTagListHandler)
		}

		items := v1.Group("/items")
		{
			items.POST("", controllor.CreateItemHandler)
			items.GET("/summary", controllor.GetItemsSummaryHandler)
			items.GET("", controllor.GetItemsHandler)
			items.GET("/balance", controllor.GetBalanceHandler)
		}
	}

	return r
}
