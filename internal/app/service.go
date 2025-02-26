package app

import (
	"t_astrum/internal/promo/handlers/http"
	"t_astrum/internal/promo/repository"
	promoUsecase "t_astrum/internal/promo/usecase"
)

func (app *App) InitService() error {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     app.cfg.RedisConfig.Host + ":" + app.cfg.RedisConfig.Port,
	//	Password: app.cfg.RedisConfig.Password,
	//	DB:       app.cfg.RedisConfig.DB,
	//})
	//
	//if err := client.Ping(context.Background()).Err(); err != nil {
	//	app.log.Info("redis init fail")
	//	return err
	//}

	//appRedis := app_redis.NewRedis(client)
	promoRepo := promoRepository.NewRepository(app.db)
	promoUC := promoUsecase.NewUsecase(promoRepo)
	promoHttp := v1.NewHandlers(promoUC)

	//_ = appRedis

	promoHttp.PromoRoutes(&app.gin.RouterGroup)

	return nil
}
