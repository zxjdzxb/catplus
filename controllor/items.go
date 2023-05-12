package controllor

import (
	"catplus-server/common"
	"catplus-server/database"
	"catplus-server/model"
	"log"
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

	// 设置账目的 HappenAt 字段为 MySQL 服务器的本地时间
	item.HappenAt = time.Now().Local()
	// 调用 BeforeSave() 方法，将 TagIDs 序列化为字符串
	item.BeforeSave()
	if item.Kind != "income" && item.Kind != "express" {
		common.Fail(c, gin.H{}, "invalid kind")
		return
	}

	if err := db.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := gin.H{
		"item": item,
	}

	item.AfterSave()

	c.JSON(http.StatusOK, response)
}

type Group struct {
	HappenAt time.Time `json:"happen_at"`
	Amount   int       `json:"amount"`
}

func GetItemsSummaryHandler(c *gin.Context) {
	// 解析查询参数
	happenedAfter := c.Query("happened_after")
	happenedBefore := c.Query("happened_before")
	kind := c.Query("kind")
	groupBy := c.Query("group_by")

	// 校验必填参数
	if happenedAfter == "" || happenedBefore == "" || kind == "" || groupBy == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required parameters"})
		return
	}

	db := database.GetDB()

	// 查询统计信息
	var groups []Group

	query := db.Table("items").
		Select(groupBy + " AS happen_at,SUM(amount) AS amount").
		Where("tag_ids IS NOT NULL")

	if happenedAfter != "" {
		query = query.Where("happen_at >= ?", happenedAfter)
	}

	if happenedBefore != "" {
		query = query.Where("happen_at <= ?", happenedBefore)
	}

	if kind != "" {
		query = query.Where("kind = ?", kind)
	}

	err := query.Group(groupBy).Scan(&groups).Error

	if err != nil {
		log.Println("Failed to fetch item summary:", err)
		return
	}
	// 计算总金额
	var total int
	err = db.Table("items").
		Select("SUM(amount) AS total").
		Where("tag_ids IS NOT NULL").
		Where("happen_at >= ?", happenedAfter).
		Where("happen_at <= ?", happenedBefore).
		Where("kind = ?", kind).
		Scan(&total).Error
	if err != nil {
		log.Println("Failed to calculate total:", err)
		return
	}

	// 构建响应数据
	response := gin.H{
		"groups": groups,
		"total":  total,
	}

	c.JSON(http.StatusOK, response)
}
