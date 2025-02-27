package DTOs

import "github.com/google/uuid"

type CreatePromocodeRequest struct {
	Code     string    `json:"code"`
	MaxUses  int       `json:"max_uses"`
	RewardId uuid.UUID `json:"reward_id"`
}

type PromocodeResponse struct {
	Code    string `json:"code"`
	MaxUses int    `json:"max_uses"`
}
