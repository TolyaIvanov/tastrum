package repository

import (
	"fmt"
	"t_astrum/internal/promo/entities"
)

// Получить всеx игроков.
func (r *Repository) GetPlayers() ([]entities.Player, error) {
	var players []entities.Player

	query := "SELECT id, username, email, created_at FROM players"
	err := r.DB.Select(&players, query)
	if err != nil {
		return nil, fmt.Errorf("error getting players from DB: %w", err)
	}

	return players, nil
}
