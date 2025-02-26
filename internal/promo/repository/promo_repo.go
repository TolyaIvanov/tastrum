package repository

import (
	"github.com/jmoiron/sqlx"
	"t_astrum/internal/promo/entities"
)

type PromocodeRepository struct {
	DB *sqlx.DB
}

func NewPromocodeRepository(db *sqlx.DB) *PromocodeRepository {
	return &PromocodeRepository{DB: db}
}

// Create сохраняет новый промокод в базу данных.
func (r *PromocodeRepository) CreatePromocode(promocode *entities.Promocode) error {
	_, err := r.DB.Exec(`
		INSERT INTO promocodes (code, max_uses, uses_count)
		VALUES ($1, $2, $3)
	`, promocode.Code, promocode.MaxUses, promocode.UsesCount)
	return err
}

// Apply применяет промокод и обновляет количество использований.
func (r *PromocodeRepository) ApplyPromocode(code string) (*entities.Promocode, error) {
	var promocode entities.Promocode
	err := r.DB.Get(&promocode, "SELECT * FROM promocodes WHERE code=$1", code)
	if err != nil {
		return nil, err
	}
	if promocode.UsesCount >= promocode.MaxUses {
		return nil, nil
	}
	_, err = r.DB.Exec("UPDATE promocodes SET uses_count = uses_count + 1 WHERE code = $1", code)
	if err != nil {
		return nil, err
	}
	return &promocode, nil
}
