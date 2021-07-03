package model

import (
	"time"
)

type Message struct {
	ID            uint64 `gorm:"AUTO_INCREMENT;primary_key;index"`
	ApplicationID uint64
	Message       string `gorm:"type:text"`
	Topic         string `gorm:"type:text"`
	Date          time.Time
}
