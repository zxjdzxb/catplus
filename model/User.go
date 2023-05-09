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
	Email              string `gorm:"not null;unique"`
	VerificationCode   string `gorm:"size:6"`
	VerificationExpiry time.Time
}
