package usecase

import (
	"t_astrum/internal/promo/entities"
)

// GetRewards юзкейс для всех наград.
func (u *Usecase) GetRewards() ([]entities.Reward, error) {
	rewards, err := u.Repository.GetRewards()
	if err != nil {
		return nil, err
	}

	return rewards, nil
}
