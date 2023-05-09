package model

import (
	"time"

	"gorm.io/gorm"
)

/* CREATE TABLE users (
		Email:              email,
		Password:           string(hashedPassword),
		VerificationCode:   code,
		VerificationExpiry: expiry,
); */

type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Email              string `gorm:"not null;unique"`
	VerificationCode   string `gorm:"size:6"`
	VerificationExpiry time.Time
}

// name
type Tag struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null;unique"`
	Sign      string `gorm:"type:varchar(20);not null;unique"`
	Kind      string `gorm:"type:varchar(20);not null;unique"`
	UserID    uint   `gorm:"primaryKey;autoIncrement"`
	DeletedAt *time.Time
}
