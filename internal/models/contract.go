package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Contract struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type ContractEvent struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	ContractID uuid.UUID `json:"contract_id"`
	EventID    uuid.UUID `json:"event_id"`
	Active     bool      `json:"active"`
	Contract   *Contract `json:"contract" gorm:"foreignKey:ContractID"`
	Event      *Event    `json:"event" gorm:"foreignKey:EventID"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

func (c *ContractEvent) TableName() string {
	return "contract_events"
}
