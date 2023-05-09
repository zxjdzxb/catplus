package controllor

import (
	"catplus-server/database"
	"catplus-server/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	//name 为 tag 的名称
	//kind 为 tag 的类型
	//sign 为 tag 的标识
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	sign := c.PostForm("sign")

	db := database.GetDB()
	var tag model.Tag

	// if err := c.ShouldBindJSON(&tag); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Validate required parameters
	if kind != "income" && kind != "expres" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid parameter(s)"})
		return
	}
	//获取用户ID
	userID := 1

	//查询数据库中是否有该name
	if err := db.Where("name = ? AND user_id = ?", name, userID).First(&tag).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已经有该标签"})
		return
	}
	db.Create(&model.Tag{Name: name, Kind: kind, Sign: sign, UserID: uint(userID)})

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"resource": gin.H{
			"id":         tag.ID,
			"name":       tag.Name,
			"sign":       tag.Sign,
			"user_id":    tag.UserID,
			"deleted_at": tag.DeletedAt,
		},
	})

}
