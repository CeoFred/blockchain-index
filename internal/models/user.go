package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type UserAction struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	EventLogID      uuid.UUID `json:"event_log_id"`
	TransactionHash string    `json:"transaction_hash"`
	EventLog        *EventLog `json:"event_log" gorm:"foreignKey:EventLogID"`
	UserID          uuid.UUID `json:"user_id"`
	Action          string    `json:"action"`
	User            *User     `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt       time.Time `json:"created_at"`
}

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Address   string    `json:"address" gorm:"uniqueIndex"`
	CreatedAt time.Time `json:"created_at"`
}
