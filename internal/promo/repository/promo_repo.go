package repository

import (
	"database/sql"
	"errors"
	"t_astrum/internal/promo/entities"
)

// Create сохраняет новый промокод в базу данных.
func (r *Repository) CreatePromocode(promocode *entities.Promocode) error {
	_, err := r.DB.Exec(`
		INSERT INTO promocodes (code, max_uses, uses_count, reward_id)
		VALUES ($1, $2, $3, $4)
	`, promocode.Code, promocode.MaxUses, promocode.UsesCount, promocode.RewardId)

	return err
}

// Apply применяет промокод и обновляет количество использований.
func (r *Repository) ApplyPromocode(code string) (*entities.Promocode, error) {
	var promocode entities.Promocode

	err := r.DB.Get(&promocode, `
		SELECT id, code, max_uses, uses_count, reward_id 
		FROM promocodes 
		WHERE code=$1
	`, code)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrPromocodeNotFound
		}
		return nil, err
	}

	if promocode.UsesCount >= promocode.MaxUses {
		return nil, entities.ErrPromocodeMaxUsesReached
	}

	_, err = r.DB.Exec("UPDATE promocodes SET uses_count = uses_count + 1 WHERE code = $1", code)
	if err != nil {
		return nil, err
	}

	return &promocode, nil
}
