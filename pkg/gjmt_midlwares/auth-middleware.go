package gjmt_midlwares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_db_service"
	"os"
	"strings"
)

func GetAuthUser(c fiber.Ctx) (*gjmt_models.User, error) {

	userID := c.Locals("userID")
	if userID == nil {
		return nil, errors.New("User ID not found in context")
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		return nil, errors.New("Invalid user ID type")
	}

	user := gjmt_models.User{}
	db := gjmt_db_service.Connect()

	if db.First(&user, userIDUint).Error != nil {
		return nil, errors.New("Error code id")
	}

	return &user, nil
}

func AuthMiddleware() fiber.Handler {

	godotenv.Load()
	var jwtKey = []byte(os.Getenv("APP_SECRET_KEY"))

	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(gjmt_models.ResponseWithError("Authorization header is required"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &gjmt_models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(gjmt_models.ResponseWithError("Invalid auth token"))
		}

		// Устанавливаем значения в `Locals` для дальнейшего использования
		c.Locals("userID", claims.Id)
		c.Locals("Username", claims.Username)

		return c.Next()
	}
}
