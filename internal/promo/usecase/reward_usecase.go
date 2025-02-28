package usecase

import (
	"t_astrum/internal/promo/entities"
)

// GetRewards юзкейс для всех наград.
func (u *RewardUsecase) GetRewards() ([]entities.Reward, error) {
	rewards, err := u.RewardRepo.GetRewards()
	if err != nil {
		return nil, err
	}

	return rewards, nil
}
