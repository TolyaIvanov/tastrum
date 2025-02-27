package DTOs

import "t_astrum/internal/promo/entities"

type RewardsResponse struct {
	Rewards []entities.Reward `json:"rewards"`
}
