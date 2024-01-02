package tools

import (
	"github.com/google/uuid"
	"time"
)

// GetUUID
// @Description get database uuid
func GetUUID() string {
	return uuid.New().String()
}

func IsValidDateTime(input string) bool {
	layout := "2006-01-02 15:04:05"
	_, err := time.Parse(layout, input)
	return err == nil
}
