package main

import (
	"catplus-server/database"
	"catplus-server/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()

	// 连接到 MySQL 数据库
	db := database.InitDb()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	// 创建 Gin 引擎
	r := gin.Default()
	r = router.CollectRoute(r)
	// 处理验证验证码的请求

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
