package controllor

import (
	"catplus-server/database"
	"catplus-server/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateItemHandler(c *gin.Context) {
	var item model.Item
	//获取mysql本地时间
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the item in the database
	db := database.GetDB()

	// 创建账目
	// 设置账目的 HappenAt 字段为 MySQL 服务器的本地时间
	item.HappenAt = time.Now().Local()

	if err := db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	repsone := gin.H{
		"item": item,
	}

	c.JSON(http.StatusOK, repsone)
}

func GetSummaryHandler(c *gin.Context) {
	// Retrieve query parameters
	happenedAfter := c.Query("happened_after")
	happenedBefore := c.Query("happened_before")
	kind := c.Query("kind")
	groupBy := c.Query("group_by")

	// Prepare the query conditions
	db := database.GetDB()
	query := db.Model(&model.Item{})

	if happenedAfter != "" {
		happenedAfterTime, _ := time.Parse(time.RFC3339, happenedAfter)
		query = query.Where("happen_at >= ?", happenedAfterTime)
	}

	if happenedBefore != "" {
		happenedBeforeTime, _ := time.Parse(time.RFC3339, happenedBefore)
		query = query.Where("happen_at <= ?", happenedBeforeTime)
	}

	if kind != "" {
		query = query.Where("kind = ?", kind)
	}

	// Perform the grouping and aggregation
	var result []model.SummaryGroup
	query = query.Select(groupBy + " AS name, SUM(amount) AS total").Group(groupBy).Scan(&result)

	c.JSON(http.StatusOK, gin.H{
		"groups": result,
	})
}
