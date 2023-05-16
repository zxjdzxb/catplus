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
	Note      string    `json:"note"`
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

func (i *Item) AfterSave() error {
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
	Group Group `json:"group" gorm:"-"`
	Total int   `json:"total"`
}

type Group struct {
	HappenAt time.Time `json:"happen_at" gorm:"column:happen_at"`
	Tag      string    `json:"tag" gorm:"column:tag"`
	Amount   int       `json:"amount" gorm:"column:amount"`
}

type Balance struct {
	Expenses int64 `json:"expenses"`
	Income   int64 `json:"income"`
}
