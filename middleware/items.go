package middleware

import (
	"catplus-server/database"
	"catplus-server/model"
	"strconv"
	"strings"
)

func GetTagsByIDs(tagIDs []int) ([]model.Tag, error) {
	var tags []model.Tag
	var db = database.GetDB()
	err := db.Where("id in (?)", tagIDs).Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func ParseTagIDs(tagIDsStr string) ([]int, error) {
	if tagIDsStr == "" {
		return nil, nil
	}

	tagIDs := strings.Split(tagIDsStr, ",")
	result := make([]int, 0, len(tagIDs))

	for _, idStr := range tagIDs {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}
		result = append(result, id)
	}

	return result, nil
}

func CalculateTotalAmount(groups []model.SummaryGroup2) int {
	total := 0
	for _, group := range groups {
		total += group.Amount
	}
	return total
}
