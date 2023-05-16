package controllor

import (
	"catplus-server/common"
	"catplus-server/database"
	"catplus-server/model"
	"log"
	"net/http"
	"strconv"
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
	if item.Kind != "income" && item.Kind != "expense" {
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
	var groups []model.Group

	query := db.Table("items").
		Select(groupBy+" AS happen_at, SUM(amount) AS amount").
		Where("tag_ids IS NOT NULL").
		Where("happen_at >= ?", happenedBefore).
		Where("happen_at <= ?", happenedAfter).
		Where("kind = ?", kind).
		Group(groupBy)

	err := query.Find(&groups).Error
	if err != nil {
		log.Println("Failed to fetch item summary:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch item summary"})
		return
	}
	// 计算总金额
	var total int
	err = db.Table("items").
		Select("SUM(amount) AS total").
		Where("tag_ids IS NOT NULL").
		Where("happen_at >= ?", happenedBefore).
		Where("happen_at <= ?", happenedAfter).
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

func GetItemsHandler(c *gin.Context) {
	page := c.Query("page")
	happenedAfter := c.Query("happened_after")
	happenedBefore := c.Query("happened_before")

	// 构建查询条件
	query := database.DB.Model(&model.Item{})

	if happenedAfter != "" {
		query = query.Where("happen_at >= ?", happenedBefore)
	}

	if happenedBefore != "" {
		query = query.Where("happen_at <= ?", happenedAfter)
	}

	// 分页查询
	var items []model.Item
	var totalCount int64
	PageSize := 10

	err := query.Count(&totalCount).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count items"})
		return
	}

	if page != "" {
		// 解析页码
		pageNum, err := strconv.Atoi(page)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		// 分页查询
		offset := (pageNum - 1) * PageSize
		query = query.Offset(offset).Limit(PageSize)
	}

	err = query.Select("id, amount,kind,happen_at,note").Find(&items).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	// 构建响应数据
	response := gin.H{
		"resources": make([]gin.H, len(items)),
		"pager": gin.H{
			"page":     page,
			"per_page": PageSize,
			"count":    totalCount,
		},
	}
	// var tags []model.Tag
	for i, item := range items {
		response["resources"].([]gin.H)[i] = gin.H{
			"id":        item.ID,
			"amount":    item.Amount,
			"kind":      item.Kind,
			"note":      item.Note,
			"happen_at": item.HappenAt,
			// "tags":      tags,
		}
	}

	c.JSON(http.StatusOK, response)
}

func GetBalanceHandler(c *gin.Context) {
	happenedAfter := c.Query("happened_after")
	happenedBefore := c.Query("happened_before")

	// 构建查询条件
	query := database.DB.Model(&model.Item{})

	if happenedAfter != "" {
		query = query.Where("happen_at >= ?", happenedAfter)
	}

	if happenedBefore != "" {
		query = query.Where("happen_at <= ?", happenedBefore)
	}

	// 计算净收入、支出和收入
	var balance int64
	var expenses int64
	var income int64

	err := query.Select("SUM(CASE WHEN kind = 'expenses' THEN amount ELSE 0 END) AS expenses").
		Select("SUM(CASE WHEN kind = 'income' THEN amount ELSE 0 END) AS income").
		Row().
		Scan(&expenses, &income)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
		return
	}

	balance = income - expenses

	// 构建响应数据
	response := gin.H{
		"balance":  balance,
		"expenses": expenses,
		"income":   income,
	}

	c.JSON(http.StatusOK, response)
}
