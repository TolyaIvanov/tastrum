package usecase

import (
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	_ "t_astrum/internal/promo/repository"
)

// ApplyPromocode применяет промокод для игрока.
func (u *Usecase) ApplyPromocode(code string) (*DTOs.PromocodeResponse, error) {
	promocode, err := u.Repository.ApplyPromocode(code)
	if err != nil {
		return nil, err
	}

	// Преобразуем в DTO перед возвратом
	return &DTOs.PromocodeResponse{
		Code:    promocode.Code,
		MaxUses: promocode.MaxUses,
	}, nil
}

// CreatePromocode создает новый промокод.
func (u *Usecase) CreatePromocode(promocode *entities.Promocode) error {
	return u.Repository.CreatePromocode(promocode)
}
