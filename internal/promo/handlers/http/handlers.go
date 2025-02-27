package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/usecase"
)

type Handlers struct {
	Usecase *usecase.Usecase
}

func NewHandlers(usecase *usecase.Usecase) *Handlers {
	return &Handlers{
		Usecase: usecase,
	}
}

func (h *Handlers) AdminPage(c *gin.Context) {
	// Получить список игроков и наград из БД
	players, err := h.Usecase.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	rewards, err := h.Usecase.GetRewards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.HTML(http.StatusOK, "admin.html", gin.H{
		"Players": players,
		"Rewards": rewards,
	})
}

// CreatePromocode создает новый промокод.
func (h *Handlers) CreatePromocode(c *gin.Context) {
	var request DTOs.CreatePromocodeRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	promocode := &entities.Promocode{
		Code:     request.Code,
		MaxUses:  request.MaxUses,
		RewardId: request.RewardId,
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

	promocode, err := h.Usecase.ApplyPromocode(code)
	if err != nil {
		if errors.Is(err, entities.ErrPromocodeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Promocode not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Преобразуем Entity в DTO
	response := DTOs.PromocodeResponse{
		Code:    promocode.Code,
		MaxUses: promocode.MaxUses,
	}

	c.JSON(http.StatusOK, response)
}

// GetPlayers возвращает всех игроков из бд.
func (h *Handlers) GetPlayers(c *gin.Context) {
	players, err := h.Usecase.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DTOs.PlayersResponse{
		Players: players,
	})
}

// GetRewards возвращает все reward из бд.
func (h *Handlers) GetRewards(c *gin.Context) {
	rewards, err := h.Usecase.GetRewards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, DTOs.RewardsResponse{
		Rewards: rewards,
	})
}
