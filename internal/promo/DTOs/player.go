package DTOs

import (
	"github.com/google/uuid"
	"time"
)

type PlayerResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type PlayersResponse struct {
	Players []PlayerResponse `json:"players"`
}
