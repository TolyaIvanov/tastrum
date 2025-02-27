package entities

import (
	"github.com/google/uuid"
	"time"
)

type Reward struct {
	Id         uuid.UUID `json:"id"`
	Reward     string    `json:"reward"`
	Created_at time.Time `json:"created_at"`
}
