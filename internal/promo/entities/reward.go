package entities

import (
	"github.com/google/uuid"
	"time"
)

type Reward struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Reward    string    `json:"reward" db:"reward"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
