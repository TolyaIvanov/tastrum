package app

import (
	"t_astrum/internal/promo/handlers/http"
	"t_astrum/internal/promo/repository"
	promoUsecase "t_astrum/internal/promo/usecase"
)

func (app *App) InitService() error {
	Repo := repository.NewRepository(app.DB)
	UC := promoUsecase.NewPromocodeUsecase(Repo)
	Http := v1.NewHandlers(UC)

	Http.PromoRoutes(&app.gin.RouterGroup)

	return nil
}
