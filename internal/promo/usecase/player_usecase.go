package usecase

import "t_astrum/internal/promo/entities"

func (u *Usecase) GetPlayers() ([]entities.Player, error) {
	players, err := u.Repository.GetPlayers()
	if err != nil {
		return nil, err
	}

	return players, nil
}
