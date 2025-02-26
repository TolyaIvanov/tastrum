package promoRepository

import (
	"github.com/jmoiron/sqlx"
	"t_astrum/internal/promo"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) promo.Repository {
	return &repository{db: db}
}
