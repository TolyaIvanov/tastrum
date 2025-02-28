package app

import (
	"t_astrum/internal/promo/handlers/http"
	"t_astrum/internal/promo/repository"
	"t_astrum/internal/promo/usecase"
)

func (app *App) InitService() error {
	Repo := repository.NewRepository(app.DB)
	promoUC := usecase.NewPromocodeUsecase(Repo)
	playerUC := usecase.NewPlayerUsecase(Repo)
	rewardUC := usecase.NewRewardUsecase(Repo)
	Http := v1.NewHandlers(promoUC, playerUC, rewardUC)

	Http.PromoRoutes(&app.gin.RouterGroup)

	return nil
}
