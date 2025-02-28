package usecase

import "t_astrum/internal/promo/entities"

func (u *PlayerUsecase) GetPlayers() ([]entities.Player, error) {
	players, err := u.PlayerRepo.GetPlayers()
	if err != nil {
		return nil, err
	}

	return players, nil
}
