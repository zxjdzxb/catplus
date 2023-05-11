package model

import (
	"time"
)

type Item struct {
	ID         uint      `gorm:"primaryKey"`
	Amount     int       `gorm:"type:int;not null"`         // amount
	Kind       string    `gorm:"type:varchar(20);not null"` // item kind
	TagIds     uint      `gorm:"type:int;not null"`         // tag id
	HappenedAt time.Time // happened_at

}
