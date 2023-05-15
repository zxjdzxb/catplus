package main

import (
	"catplus-server/config"
	"catplus-server/database"
	"catplus-server/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	// 连接到 MySQL 数据库
	db := database.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 创建 Gin 引擎
	r := gin.Default()
	store := cookie.NewStore([]byte("secret")) // Replace "secret" with your desired session secret
	r.Use(sessions.Sessions("mySession", store))
	r = router.CollectRoute(r)
	// 处理验证验证码的请求

	// 运行 Gin 应用程序
	if err := r.Run(":8080"); err != nil {
		panic("failed to start server")
	}
}
