package gjmt_db_service

import (
	"fmt"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
)

func Migrate() {

	fmt.Println("[Migrate] Migratation ")
	db := Connect()

	db.AutoMigrate(&gjmt_models.User{})

}
