package pkg

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

// Функция генерации рандомной строки заданной длины
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := strings.Builder{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < length; i++ {
		index := rand.Intn(len(charset))
		result.WriteByte(charset[index])
	}

	return result.String()
}

// Основная функция, которая добавляет хеш текущей даты и времени
func GenerateRandomStringWithDateHash() string {
	// Генерация рандомной строки
	randomString := GenerateRandomString(16)

	// Получение текущей даты и времени
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Хеширование текущей даты и времени
	hash := sha256.Sum256([]byte(currentTime))

	// Конвертация хеша в строку (hex)
	hashString := hex.EncodeToString(hash[:])

	// Объединяем рандомную строку и хеш
	return randomString + "_" + hashString
}
