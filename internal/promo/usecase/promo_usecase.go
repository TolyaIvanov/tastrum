package promoUsecase

import "t_astrum/internal/promo"

type Usecase struct {
	repo promo.Repository
}

func NewUsecase(repo promo.Repository) *Usecase {
	return &Usecase{repo: repo}
}
