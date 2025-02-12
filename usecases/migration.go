package usecases

import (
	"auth-service/models"
	"auth-service/pkg/db_service"
	"fmt"
)

func Migrate() {

	fmt.Println("[Migrate] Migratation ")
	db := db_service.Connect()

	db.AutoMigrate(&models.User{})

}
