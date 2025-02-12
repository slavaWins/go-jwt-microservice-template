package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID         uint `gorm:"primarykey" json:"id"`
	Created_at time.Time
	//updated_at time.Time
	Email string
}

type Claims struct {
	Username string
	Id       uint
	jwt.StandardClaims
}
