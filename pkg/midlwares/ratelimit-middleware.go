package midlwares

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"
	"time"
)

func RateLimitMiddleware(requestsPerHour int) gin.HandlerFunc {
	// Создаем лимитер с указанным количеством запросов в час
	rate := limiter.Rate{
		//Period: time.Hour,
		Period: time.Minute,
		Limit:  int64(requestsPerHour),
	}

	// Используем memory store для хранения состояния лимитера
	store := memory.NewStore()

	// Создаем лимитер
	limiterInstance := limiter.New(store, rate)

	return func(c *gin.Context) {
		// Получаем IP адрес клиента
		ip := c.ClientIP()

		// Проверяем лимит для данного IP
		context, err := limiterInstance.Get(c, ip)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		if context.Reached {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Превышено число попыток авторизации"})
			c.Abort()
			return
		}

		// Продолжаем выполнение запроса
		c.Next()
	}
}
