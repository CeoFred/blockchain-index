package models

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Signature   string    `json:"signature" gorm:"uniqueIndex"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (e *Event) TableName() string {
	return "events"
}

type EventLog struct {
	ID              uuid.UUID       `json:"id" gorm:"primaryKey"`
	ContractEventID uuid.UUID       `json:"contract_event_id"`
	ContractID      uuid.UUID       `json:"contract_id"`
	ContractAddress string          `json:"contract_address"`
	EventName       string          `json:"event_name"`
	BlockNumber     uint64          `json:"block_number"`
	TransactionHash string          `json:"transaction_hash"`
	LogIndex        uint            `json:"log_index"`
	Data            json.RawMessage `json:"data" gorm:"type:jsonb"`
	Topics          json.RawMessage `json:"topics" gorm:"type:jsonb"`
	Timestamp       time.Time       `json:"timestamp"`
	CreatedAt       time.Time       `json:"created_at"`
}

func (e *EventLog) TableName() string {
	return "event_logs"
}
