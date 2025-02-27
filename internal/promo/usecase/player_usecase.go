package usecase

import "t_astrum/internal/promo/entities"

// GetRewards юзкейс для всех players.
func (u *Usecase) GetPlayers() ([]entities.Player, error) {
	return u.Repository.GetPlayers()
}
