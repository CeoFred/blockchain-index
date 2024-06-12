package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Event struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Signature string `json:"signature"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
}


func (e *Event) TableName() string {
	return "events"
}

type EventLog struct {
	ID uuid.UUID `json:"id"`
	ContractAddress uuid.UUID `json:"contract_address"`
	EventName string `json:"event_name"`
	BlockNumber int `json:"block_number"`
	TransactionHash uuid.UUID `json:"transaction_hash"`
	LogIndex uint `json:"log_index"`
	Data  map[string]interface{}  `json:"data"`
	Topics []string `json:"topics"`
	Timestamp time.Time `json:"timestamp"`
}

func (e *EventLog) TableName() string {
	return "event_logs"
}