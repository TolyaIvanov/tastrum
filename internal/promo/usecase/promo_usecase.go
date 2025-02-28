package usecase

import (
	"log"
	"t_astrum/internal/promo/entities"
	_ "t_astrum/internal/promo/repository"
)

// ApplyPromocode применяет промокод для игрока.
func (u *PromocodeUsecase) ApplyPromocode(code string) (*entities.Promocode, error) {
	promocode, err := u.PromoRepo.ApplyPromocode(code)
	if err != nil {
		return nil, err
	}

	return promocode, nil
}

// CreatePromocode создает новый промокод.
func (u *PromocodeUsecase) CreatePromocode(promocode *entities.Promocode) error {
	log.Println("CreatePromocode: received request", promocode)

	exists, err := u.PromoRepo.PromocodeExists(promocode.Code)
	if err != nil {
		log.Println("Error in CreatePromocode:", err)
		return err
	}
	if exists {
		return entities.ErrPromocodeExists
	}

	log.Println("CreatePromocode: created promo", promocode)
	return u.PromoRepo.CreatePromocode(promocode)
}
