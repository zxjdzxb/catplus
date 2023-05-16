package middleware

import (
	"catplus-server/database"
	"catplus-server/model"
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
