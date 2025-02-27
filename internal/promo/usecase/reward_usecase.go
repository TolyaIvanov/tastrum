package usecase

import (
	"t_astrum/internal/promo/entities"
)

// GetRewards юзкейс для всех наград.
func (u *Usecase) GetRewards() ([]entities.Reward, error) {
	return u.Repository.GetRewards()
}
