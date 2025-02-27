package repository

import (
	"fmt"
	"t_astrum/internal/promo/entities"
)

// Получить все награды.
func (r *Repository) GetRewards() ([]entities.Reward, error) {
	var rewards []entities.Reward

	query := fmt.Sprintf("SELECT * FROM rewards;")
	err := r.DB.Select(&rewards, query)

	return rewards, err
}
