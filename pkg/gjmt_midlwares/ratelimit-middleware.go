package gjmt_midlwares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"
	"time"
)

func RateLimitMiddleware(requestsPerHour int) fiber.Handler {
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

	return func(c fiber.Ctx) error {
		// Получаем IP адрес клиента
		ip := c.IP()

		// Проверяем лимит для данного IP
		context, err := limiterInstance.Get(c.Context(), ip)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(gjmt_models.ResponseWithError("Internal server error"))
		}

		if context.Reached {
			return c.Status(http.StatusTooManyRequests).JSON(gjmt_models.ResponseWithError("Превышено число попыток авторизации"))
		}

		// Продолжаем выполнение запроса
		return c.Next()

	}
}
