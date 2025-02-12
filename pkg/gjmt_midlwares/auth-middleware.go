package gjmt_midlwares

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_db_service"
	"net/http"
	"os"
	"strings"
)

func GetAuthUser(c *gin.Context) (*gjmt_models.User, error) {

	userID, exists := c.Get("userID")
	if !exists {
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

func AuthMiddleware() gin.HandlerFunc {

	godotenv.Load()
	var jwtKey = []byte(os.Getenv("APP_SECRET_KEY"))

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {

			c.JSON(http.StatusUnauthorized, gjmt_models.NewErrorResponse("Authorization header is required"))
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &gjmt_models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gjmt_models.NewErrorResponse("Invalid auth token"))
			c.Abort()
			return
		}

		c.Set("userID", claims.Id)
		c.Set("Username", claims.Username)

		c.Next()
	}
}
