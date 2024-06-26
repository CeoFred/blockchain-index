package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type FaucetAsset string

const (
	FaucetAssetRWA  FaucetAsset = "RWA"
	FaucetAssetUSDT FaucetAsset = "USDT"
)

type Faucet struct {
	ID         uuid.UUID   `json:"id" gorm:"primaryKey"`
	Asset      FaucetAsset `json:"asset"`
	Address    string      `json:"address"`
	StartBlock uint        `json:"start_block"`
	EndBlock   uint        `json:"end_block"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at" default:"null"`
}

type FaucetTransfer struct {
	ID              uuid.UUID `json:"id" gorm:"primaryKey"`
	FaucetID        uuid.UUID `json:"faucet_id"`
	Recipient       string    `json:"recipient"`
	Value           uint      `json:"value"`
	TransactionHash string    `json:"transaction_hash"`
	BlockNumber     string    `json:"block_number"`
	Timestamp       time.Time `json:"timestamp"`
	CreatedAt       time.Time `json:"created_at"`
}
