package usecases

import (
	"fmt"
	"github.com/slavaWins/go-jwt-microservice-template/models"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/db_service"
)

func Migrate() {

	fmt.Println("[Migrate] Migratation ")
	db := db_service.Connect()

	db.AutoMigrate(&models.User{})

}
