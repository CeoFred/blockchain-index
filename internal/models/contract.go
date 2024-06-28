package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type ContractEventStatus bool
type ContractCategory string

const (
	ContractEventStatusActive   ContractEventStatus = true
	ContractEventStatusInactive ContractEventStatus = false

	ContractCategorySwap   ContractCategory = "swap"
	ContractCategoryERC20  ContractCategory = "erc20"
	ContractCategoryBridge ContractCategory = "bridge"
)

type Contract struct {
	ID         uuid.UUID        `json:"id" gorm:"primaryKey"`
	Name       string           `json:"name"`
	Address    string           `json:"address"`
	StartBlock uint             `json:"start_block"`
	EndBlock   uint             `json:"end_block"`
	Category   ContractCategory `json:"category" default:"erc20"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
	DeletedAt  time.Time        `json:"deleted_at"`
}

type ContractEvent struct {
	ID         uuid.UUID `json:"id" gorm:"primaryKey"`
	ContractID uuid.UUID `json:"contract_id" gorm:"uniqueIndex:idx_contract_event"`
	EventID    uuid.UUID `json:"event_id" gorm:"uniqueIndex:idx_contract_event"`

	Active    ContractEventStatus `json:"active"`
	Contract  *Contract           `json:"contract" gorm:"foreignKey:ContractID"`
	Event     *Event              `json:"event" gorm:"foreignKey:EventID"`
	CreatedAt time.Time           `json:"created_at"`
	DeletedAt time.Time           `json:"deleted_at"`
}

func (c *ContractEvent) TableName() string {
	return "contract_events"
}
