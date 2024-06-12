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