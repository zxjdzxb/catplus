package model

import "time"

/* CREATE TABLE users (
		Email:              email,
		Password:           string(hashedPassword),
		VerificationCode:   code,
		VerificationExpiry: expiry,
); */

type User struct {
	ID                 uint   `gorm:"primaryKey"`
	Email              string `gorm:"unique"`
	Password           string
	VerificationCode   string
	VerificationExpiry time.Time
}
