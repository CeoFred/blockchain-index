package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Contract struct {
	ID        uuid.UUID `json:"id" validate:"required"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
