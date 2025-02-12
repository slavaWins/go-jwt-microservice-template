package gjmt_models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// User Базова модель пользователя
// @Description Может бысть расширена в других сервисах
type User struct {
	ID         uint `gorm:"primarykey" json:"id"`
	Created_at time.Time
	//updated_at time.Time
	Email string
}

// Claims Нагрузка в jwt токене
// @Description Может бысть расширена в других сервисах
type Claims struct {
	Username string
	Id       uint
	jwt.StandardClaims
}
