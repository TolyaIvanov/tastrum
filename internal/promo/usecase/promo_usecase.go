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

	return &DTOs.PromocodeResponse{
		Code:    promocode.Code,
		MaxUses: promocode.MaxUses,
	}, nil
}

// CreatePromocode создает новый промокод.
func (u *Usecase) CreatePromocode(promocode *entities.Promocode) error {
	exists, err := u.Repository.PromocodeExists(promocode.Code)
	if err != nil {
		return err
	}
	if exists {
		return entities.ErrPromocodeExists
	}

	return u.Repository.CreatePromocode(promocode)
}
