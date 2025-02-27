package repository

import (
	"fmt"
	"t_astrum/internal/promo/entities"
)

// Получить всеx игроков.
func (r *Repository) GetPlayers() ([]entities.Player, error) {
	var players []entities.Player

	query := fmt.Sprintf("SELECT * FROM players;")
	err := r.DB.Select(&players, query)

	return players, err
}
