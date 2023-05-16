package controllor

import (
	"catplus-server/common"
	"catplus-server/database"
	"catplus-server/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
Content-Type: application/json
{
	"name": "早餐",
	"kind": "expense",
	"sign": "-", // 只有 kind 为 expense 时才有 sign 字段
}
*/

func CreateTagHandler(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		common.Fail(c, gin.H{"error": err.Error()}, "invalid request body")
		return
	}

	// Validate name field
	if len(tag.Name) < 2 || len(tag.Name) > 6 {
		common.Fail(c, gin.H{}, "invalid name")
		return
	}

	// Validate kind field
	if tag.Kind != "income" && tag.Kind != "expense" {
		common.Fail(c, gin.H{}, "invalid kind")
		return
	}
	if tag.Kind == "income" && tag.Sign != "+" {
		common.Fail(c, gin.H{}, "invalid sign for income tag")
		return
	}

	if tag.Kind == "expense" && tag.Sign != "-" {
		common.Fail(c, gin.H{}, "invalid sign for expense tag")
		return
	}

	db := database.GetDB()

	// TODO: 从 session 中获取用户 ID
	session := sessions.Default(c)

	// Retrieve the session value
	userID := session.Get("userID")
	log.Println("userID:", userID)

	tag.UserID = userID.(uint)
	// 检查 Name 和 UserID 是否已存在
	var count int64
	db.Model(&model.Tag{}).Where("name = ? AND user_id = ?", tag.Name, tag.UserID).Count(&count)
	if count > 0 {
		common.Fail(c, gin.H{"code": http.StatusConflict}, "已有该标签")
		return
	}

	result := db.Create(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	common.Success(c, gin.H{
		"id":         tag.ID,
		"name":       tag.Name,
		"sign":       tag.Sign,
		"user_id":    tag.UserID,
		"deleted_at": tag.DeletedAt,
	}, "create tag success")
}

func UpdateTagHandler(c *gin.Context) {
	// 获取标签ID
	tagID := c.Param("id")

	// 从数据库中查找标签
	var tag model.Tag
	db := database.GetDB()
	if err := db.First(&tag, tagID).Error; err != nil {
		common.Fail(c, gin.H{}, "tag not found")
		return
	}

	// 解析请求参数
	var requestBody struct {
		Name string `json:"name"`
		Sign string `json:"sign"`
	}

	// sign 为-时，kind 必须为 expense;为+为 income

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		common.Fail(c, gin.H{"error": err.Error()}, "invalid request body")
		return
	}

	// 更新标签字段
	if requestBody.Name != "" {
		tag.Name = requestBody.Name
	}
	if requestBody.Sign != "" {
		tag.Sign = requestBody.Sign
	}

	// 保存更新后的标签到数据库
	if err := db.Save(&tag).Error; err != nil {
		common.Fail(c, gin.H{}, "failed to update tag")
		return
	}

	// 返回更新后的标签
	common.Success(c, gin.H{
		"id":         tag.ID,
		"name":       tag.Name,
		"sign":       tag.Sign,
		"user_id":    tag.UserID,
		"deleted_at": tag.DeletedAt,
	}, "update tag success")

}

func GetTagHandler(c *gin.Context) {
	// 获取标签ID
	tagID := c.Param("id")

	// 从数据库中查找标签
	var tag model.Tag
	db := database.GetDB()
	if err := db.First(&tag, tagID).Error; err != nil {
		common.Fail(c, gin.H{}, "tag not found")
		return
	}
	common.Success(c, gin.H{
		"id":         tag.ID,
		"name":       tag.Name,
		"sign":       tag.Sign,
		"user_id":    tag.UserID,
		"deleted_at": tag.DeletedAt,
	}, "get tag success")
}

func DeleteTagHandler(c *gin.Context) {
	// 获取标签ID
	tagID := c.Param("id")

	// 从数据库中查找标签
	var tag model.Tag
	db := database.GetDB()
	if err := db.First(&tag, tagID).Error; err != nil {
		common.Fail(c, gin.H{}, "tag not found")
		return
	}

	// 删除标签
	if err := db.Delete(&tag).Error; err != nil {
		common.Fail(c, gin.H{}, "failed to delete tag")
		return
	}

	common.Success(c, gin.H{}, "delete tag success")
}

/*
localhost:8080/api/v1/tags?page=1&kind=expense
*/
func GetTagListHandler(c *gin.Context) {
	// 获取查询参数
	kind := c.Query("kind") // 收入或支出
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	// 分页处理
	perPage := 10                  // 每页显示的标签数量
	offset := (page - 1) * perPage // 计算偏移量

	// 构建查询条件
	db := database.GetDB()
	query := db.Model(&model.Tag{}).
		Offset(offset).
		Limit(perPage)

	// 添加条件判断
	if kind != "" {
		query = query.Where("kind = ?", kind)
	}
	// 查询标签列表
	var tags []model.Tag
	if err := query.Find(&tags).Error; err != nil {
		common.Fail(c, gin.H{}, "failed to get tag list")
		return
	}

	// 构建响应数据
	var resources []gin.H
	for _, tag := range tags {
		resources = append(resources, gin.H{
			"id":         tag.ID,
			"name":       tag.Name,
			"sign":       tag.Sign,
			"user_id":    tag.UserID,
			"deleted_at": tag.DeletedAt,
		})
	}

	common.Success(c, gin.H{"resources": resources}, "get tag list success")
}
