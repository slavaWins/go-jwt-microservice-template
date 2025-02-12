package usecases

import (
	"fmt"
	"go-jwt-microservice-template/models"
	"go-jwt-microservice-template/pkg/db_service"
)

func Migrate() {

	fmt.Println("[Migrate] Migratation ")
	db := db_service.Connect()

	db.AutoMigrate(&models.User{})

}
