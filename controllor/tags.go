package controllor

import (
	"catplus-server/database"
	"catplus-server/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Content-Type: application/json
{
	"name": "早餐",
	"kind": "express",
	"sign": "-", // 只有 kind 为 express 时才有 sign 字段
}
*/

func CreateUserHandler(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBind(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Received request with parameters: \n%+v\t\n", tag)

	// Validate name field
	if len(tag.Name) < 2 || len(tag.Name) > 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid name"})
		return
	}

	// Validate kind field
	if tag.Kind != "income" && tag.Kind != "express" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid kind"})
		return
	}
	if tag.Kind == "income" && tag.Sign != "+" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sign for income tag"})
		return
	} else if tag.Kind == "express" && tag.Sign != "-" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sign for express tag"})
		return
	}

	db := database.GetDB()

	tag.UserID = 1 // TODO: 从 session 中获取用户 ID
	// 检查 Name 和 UserID 是否已存在
	var count int64
	db.Model(&model.Tag{}).Where("name = ? AND user_id = ?", tag.Name, tag.UserID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Tag with same name and userID already exists"})
		return
	}

	result := db.Create(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

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
