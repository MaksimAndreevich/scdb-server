package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// generateDemoImageID генерирует уникальный ID для демо-образа
func GenerateDemoImageID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return "demo-" + hex.EncodeToString(bytes)
}
