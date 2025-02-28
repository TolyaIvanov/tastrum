package repository

import (
	"github.com/jmoiron/sqlx"
	"t_astrum/internal/promo/entities"
)

type PlayerRepositoryInterface interface {
	GetPlayers() ([]entities.Player, error)
}

type RewardRepositoryInterface interface {
	GetRewards() ([]entities.Reward, error)
}

type PromoRepositoryInterface interface {
	ApplyPromocode(code string) (*entities.Promocode, error)
	CreatePromocode(promocode *entities.Promocode) error
	PromocodeExists(code string) (bool, error)
}

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{DB: db}
}
