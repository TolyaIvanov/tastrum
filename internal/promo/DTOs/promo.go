package DTOs

import "github.com/google/uuid"

type CreatePromocodeRequest struct {
	Code     string    `json:"code" binding:"required"`
	MaxUses  int       `json:"max_uses" binding:"required,gt=0"`
	RewardId uuid.UUID `json:"reward_id" binding:"required"`
}

type PromocodeResponse struct {
	Code        string `json:"code"`
	CurrentUses int    `json:"current_uses"`
	MaxUses     int    `json:"max_uses"`
}
