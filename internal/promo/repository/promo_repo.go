package repository

import (
	"database/sql"
	"errors"
	"log"
	"t_astrum/internal/promo/entities"
)

func (r *Repository) PromocodeExists(code string) (bool, error) {
	var count int

	err := r.DB.Get(&count, "SELECT COUNT(*) FROM promocodes WHERE code=$1", code)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *Repository) CreatePromocode(promocode *entities.Promocode) error {
	log.Println("CreatePromocode: inserting into DB", promocode)

	_, err := r.DB.Exec(`
		INSERT INTO promocodes (id, code, max_uses, uses_count, reward_id)
		VALUES ($1, $2, $3, $4, $5)
	`, promocode.ID, promocode.Code, promocode.MaxUses, promocode.UsesCount, promocode.RewardId)
	if err != nil {
		return err
	}

	if err != nil {
		log.Println("Error inserting promocode:", err)
		return err
	}

	log.Println("CreatePromocode: successfully inserted", promocode)
	return err
}

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

	//
	//сейчас удалил из миграции таблицу, но тут также закидываем приминение наград в зависимости от uuid игрока
	//

	return &promocode, nil
}
