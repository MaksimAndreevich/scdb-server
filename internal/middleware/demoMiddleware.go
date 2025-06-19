package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/scdb/server/internal/config"
)

// DemoLimiter хранит информацию о лимитах для демо-образа
type DemoLimiter struct {
	ImageID      string    `json:"image_id"`
	MaxRequests  int       `json:"max_requests"`
	UsedRequests int       `json:"used_requests"`
	CreatedAt    time.Time `json:"created_at"`
	mu           sync.RWMutex
}

var (
	demoLimiter *DemoLimiter
	once        sync.Once
)

const demoStatsFile = "./tmp/demo_stats.json"

// InitDemoLimiter инициализирует лимитер для демо-образа
func InitDemoLimiter() {
	once.Do(func() {
		// Пытаемся загрузить существующие данные
		if data, err := loadDemoStats(); err == nil {
			demoLimiter = data
		} else {
			// Создаем новые данные
			demoLimiter = &DemoLimiter{
				ImageID:      config.AppConfig.DemoImageID,
				MaxRequests:  50,
				UsedRequests: 0,
				CreatedAt:    time.Now(),
			}
			// Сохраняем новые данные
			saveDemoStats()
		}
	})
}

// loadDemoStats загружает статистику из файла
func loadDemoStats() (*DemoLimiter, error) {
	data, err := os.ReadFile(demoStatsFile)
	if err != nil {
		return nil, err
	}

	var limiter DemoLimiter
	if err := json.Unmarshal(data, &limiter); err != nil {
		return nil, err
	}

	return &limiter, nil
}

// saveDemoStats сохраняет статистику в файл
func saveDemoStats() error {
	demoLimiter.mu.RLock()
	defer demoLimiter.mu.RUnlock()

	data, err := json.Marshal(demoLimiter)
	if err != nil {
		return err
	}

	return os.WriteFile(demoStatsFile, data, 0644)
}

// DemoLimitMiddleware ограничивает количество запросов для демо-образа
func DemoLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Если демо-режим не включен, пропускаем
		if !config.AppConfig.DemoMode {
			c.Next()
			return
		}

		demoLimiter.mu.Lock()
		defer demoLimiter.mu.Unlock()

		// Проверяем лимит
		if demoLimiter.UsedRequests >= demoLimiter.MaxRequests {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":         "Лимит запросов исчерпан",
				"message":       "Демо-версия ограничена 50 запросами. Обратитесь к разработчику для получения нового доступа.",
				"image_id":      demoLimiter.ImageID,
				"used_requests": demoLimiter.UsedRequests,
				"max_requests":  demoLimiter.MaxRequests,
			})
			return
		}

		// Увеличиваем счетчик
		demoLimiter.UsedRequests++

		// Сохраняем в файл
		go saveDemoStats()

		// Добавляем информацию о лимитах в заголовки
		c.Header("X-Demo-Image-ID", demoLimiter.ImageID)
		c.Header("X-Demo-Requests-Used", strconv.Itoa(demoLimiter.UsedRequests))
		c.Header("X-Demo-Requests-Limit", strconv.Itoa(demoLimiter.MaxRequests))
		c.Header("X-Demo-Requests-Remaining", strconv.Itoa(demoLimiter.MaxRequests-demoLimiter.UsedRequests))

		c.Next()
	}
}

// GetDemoStats возвращает статистику использования демо
func GetDemoStats() map[string]interface{} {
	demoLimiter.mu.RLock()
	defer demoLimiter.mu.RUnlock()

	return map[string]interface{}{
		"image_id":           demoLimiter.ImageID,
		"used_requests":      demoLimiter.UsedRequests,
		"max_requests":       demoLimiter.MaxRequests,
		"remaining_requests": demoLimiter.MaxRequests - demoLimiter.UsedRequests,
		"created_at":         demoLimiter.CreatedAt,
		"is_demo":            config.AppConfig.DemoMode,
	}
}

// ResetDemoStats сбрасывает счетчик (для тестирования)
func ResetDemoStats() error {
	demoLimiter.mu.Lock()
	defer demoLimiter.mu.Unlock()

	demoLimiter.UsedRequests = 0
	demoLimiter.CreatedAt = time.Now()
	demoLimiter.ImageID = config.AppConfig.DemoImageID

	return saveDemoStats()
}
