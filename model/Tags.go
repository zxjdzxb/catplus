package model

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"` // tag name
	Sign      string `gorm:"type:varchar(20);not null"` // tag sign
	Kind      string `gorm:"type:varchar(20);not null"` // tag kind
	UserID    uint
	DeletedAt *time.Time
}
