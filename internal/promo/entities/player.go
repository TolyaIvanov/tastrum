package entities

import (
	"github.com/google/uuid"
	"time"
)

type Player struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}
