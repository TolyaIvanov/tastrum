package DTOs

import (
	"github.com/google/uuid"
	"time"
)

type RewardResponse struct {
	ID        uuid.UUID `json:"id"`
	Reward    string    `json:"reward"`
	CreatedAt time.Time `json:"created_at"`
}

type RewardsResponse struct {
	Rewards []RewardResponse `json:"rewards"`
}
