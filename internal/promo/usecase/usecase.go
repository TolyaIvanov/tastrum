package usecase

import (
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/repository"
)

type PromocodeUsecaseInterface interface {
	CreatePromocode(promocode *entities.Promocode) error
	ApplyPromocode(code string) (*entities.Promocode, error)
}

type PlayerUsecaseInterface interface {
	GetPlayers() ([]entities.Player, error)
}

type RewardUsecaseInterface interface {
	GetRewards() ([]entities.Reward, error)
}

type PromocodeUsecase struct {
	PromoRepo repository.PromoRepositoryInterface
}

func NewPromocodeUsecase(repo repository.PromoRepositoryInterface) *PromocodeUsecase {
	return &PromocodeUsecase{PromoRepo: repo}
}

type RewardUsecase struct {
	RewardRepo repository.RewardRepositoryInterface
}

func NewRewardUsecase(repo repository.RewardRepositoryInterface) *RewardUsecase {
	return &RewardUsecase{RewardRepo: repo}
}

type PlayerUsecase struct {
	PlayerRepo repository.PlayerRepositoryInterface
}

func NewPlayerUsecase(repo repository.PlayerRepositoryInterface) *PlayerUsecase {
	return &PlayerUsecase{PlayerRepo: repo}
}
