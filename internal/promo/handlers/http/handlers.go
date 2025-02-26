package v1

import (
	"t_astrum/internal/promo"
)

type Handlers struct {
	usecase promo.Usecase
}

func NewHandlers(usecase promo.Usecase) *Handlers {
	return &Handlers{
		usecase: usecase,
	}
}
