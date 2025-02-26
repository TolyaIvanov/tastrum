package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	"t_astrum/internal/promo/handlers/http"
	"t_astrum/internal/promo/repository"
	promoUsecase "t_astrum/internal/promo/usecase"
)

func (app *App) InitService() error {
	client := redis.NewClient(&redis.Options{
		Addr:     app.cfg.RedisConfig.Host + ":" + app.cfg.RedisConfig.Port,
		Password: app.cfg.RedisConfig.Password,
		DB:       app.cfg.RedisConfig.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		app.log.Info("redis init fail")
		return err
	}

	promoRepo := repository.NewPromocodeRepository(app.DB)
	promoUC := promoUsecase.NewPromocodeUsecase(promoRepo, client)
	promoHttp := v1.NewHandlers(promoUC)

	promoHttp.PromoRoutes(&app.gin.RouterGroup)

	return nil
}
