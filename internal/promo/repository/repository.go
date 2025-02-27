package repository

import (
	"github.com/jmoiron/sqlx"
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
)

type RepositoryInterface interface {
	ApplyPromocode(code string) (*DTOs.PromocodeResponse, error)
	CreatePromocode(promocode *entities.Promocode) error
	GetPlayers() ([]entities.Player, error)
	GetRewards() ([]entities.Reward, error)
}

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}
