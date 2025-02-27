package usecase

import "t_astrum/internal/promo/repository"

type Usecase struct {
	Repository *repository.Repository
}

func NewPromocodeUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		Repository: repo,
	}
}
