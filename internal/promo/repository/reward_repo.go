package repository

import (
	"fmt"
	"t_astrum/internal/promo/entities"
)

// Получить все награды.
func (r *Repository) GetRewards() ([]entities.Reward, error) {
	var rewards []entities.Reward

	query := "SELECT id, reward, created_at FROM rewards"
	err := r.DB.Select(&rewards, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching rewards: %w", err)
	}

	return rewards, nil
}
