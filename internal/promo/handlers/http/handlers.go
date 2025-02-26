package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/usecase"
)

type Handlers struct {
	Usecase *usecase.PromocodeUsecase
}

func NewHandlers(usecase *usecase.PromocodeUsecase) *Handlers {
	return &Handlers{
		Usecase: usecase,
	}
}

// CreatePromocode создает новый промокод.
func (h *Handlers) CreatePromocode(c *gin.Context) {
	var request DTOs.CreatePromocodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//todo:добавить проверку на наличие всех параметров

	promocode := &entities.Promocode{
		Code:    request.Code,
		MaxUses: request.MaxUses,
	}

	err := h.Usecase.CreatePromocode(promocode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := DTOs.PromocodeResponse{
		Code:    promocode.Code,
		MaxUses: promocode.MaxUses,
	}

	c.JSON(http.StatusOK, response)
}

// ApplyPromocode применяет промокод и возвращает результат.
func (h *Handlers) ApplyPromocode(c *gin.Context) {
	code := c.Param("code")
	promocode, err := h.Usecase.ApplyPromocode(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if promocode == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Promocode not found or expired"})
		return
	}

	c.JSON(http.StatusOK, DTOs.PromocodeResponse{
		Code: promocode.Code,
	})
}
