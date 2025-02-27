package usecase

import (
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/repository"
)

type UsecaseInterface interface {
	CreatePromocode(promocode *entities.Promocode) error
	ApplyPromocode(code string) (*DTOs.PromocodeResponse, error)
	GetPlayers() ([]entities.Player, error)
	GetRewards() ([]entities.Reward, error)
}

type Usecase struct {
	Repository *repository.Repository
}

func NewPromocodeUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Repository: repo,
	}
}
