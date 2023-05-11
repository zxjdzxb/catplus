package model

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	gorm.Model
	Amount    int       `json:"amount"`
	Kind      string    `json:"kind"`
	HappenAt  time.Time `json:"happen_at"`
	TagIDs    []int     `gorm:"-" json:"tag_ids"`
	TagIDsStr string    `gorm:"column:tag_ids" json:"-"`
}

func (i *Item) BeforeSave() error {
	// 将 TagIDs 序列化为字符串
	if i.TagIDs != nil {
		tagIDsBytes, err := json.Marshal(i.TagIDs)
		if err != nil {
			return err
		}
		i.TagIDsStr = string(tagIDsBytes)
	}
	return nil
}

func (i *Item) AfterFind() error {
	// 将字符串反序列化为 TagIDs 切片
	if i.TagIDsStr != "" {
		err := json.Unmarshal([]byte(i.TagIDsStr), &i.TagIDs)
		if err != nil {
			return err
		}
	}
	return nil
}

type SummaryGroup struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
}
