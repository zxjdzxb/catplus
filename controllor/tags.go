package controllor

import (
	"catplus-server/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name   string
	Kind   string
	Sign   string
	UserID uint
}

func CreateTag(c *gin.Context) {
	name := c.PostForm("name")
	kind := c.PostForm("kind")
	sign := c.PostForm("sign")

	var tag Tag

	// Validate required parameters
	if kind != "income" && kind != "expres" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing or invalid parameter(s)"})
		return
	}

	userID := tag.UserID
	db := database.GetDB()

	// Check for duplicate tag for a user
	if err := db.Where("name = ? AND user_id = ?", name, userID).First(&tag).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该用户已经有该标签"})
		return
	}

	// Create tag
	tag = Tag{Name: name, Kind: kind, Sign: sign, UserID: uint(userID)}
	if err := db.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create tag"})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// Check for duplicate tag for a user
