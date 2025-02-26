package usecase

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/repository"
)

type PromocodeUsecase struct {
	Repository *repository.PromocodeRepository
	Redis      *redis.Client
}

func NewPromocodeUsecase(repo *repository.PromocodeRepository, redis *redis.Client) *PromocodeUsecase {
	return &PromocodeUsecase{
		Repository: repo,
		Redis:      redis,
	}
}

// ApplyPromocode применяет промокод для игрока.
func (u *PromocodeUsecase) ApplyPromocode(ctx context.Context, code string) (*entities.Promocode, error) {
	// Проверка в Redis
	val, err := u.Redis.Get(ctx, code).Result()
	if err == redis.Nil {
		// Промокод не найден в кэше, получаем из базы
		promocode, err := u.Repository.ApplyPromocode(code)
		if err != nil {
			return nil, err
		}

		if promocode != nil {
			// Сохраняем в Redis на 1 час
			u.Redis.Set(ctx, code, promocode.UsesCount, 3600)
		}

		return promocode, nil
	}

	useCount, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	// Промокод найден в кэше
	return &entities.Promocode{
		Code:      code,
		UsesCount: useCount,
	}, nil
}

// CreatePromocode создает новый промокод.
func (u *PromocodeUsecase) CreatePromocode(promocode *entities.Promocode) error {
	return u.Repository.CreatePromocode(promocode)
}
