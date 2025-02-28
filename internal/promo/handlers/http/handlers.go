package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"t_astrum/internal/promo/DTOs"
	"t_astrum/internal/promo/entities"
	"t_astrum/internal/promo/usecase"
)

type Handlers struct {
	PromoUC  usecase.PromocodeUsecaseInterface
	PlayerUC usecase.PlayerUsecaseInterface
	RewardUC usecase.RewardUsecaseInterface
}

// Конструктор
func NewHandlers(promoUC usecase.PromocodeUsecaseInterface, playerUC usecase.PlayerUsecaseInterface, rewardUC usecase.RewardUsecaseInterface) *Handlers {
	return &Handlers{
		PromoUC:  promoUC,
		PlayerUC: playerUC,
		RewardUC: rewardUC,
	}
}

func (h *Handlers) AdminPage(c *gin.Context) {
	players, err := h.PlayerUC.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	rewards, err := h.RewardUC.GetRewards()
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
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	promocode := entities.NewPromocode(request.Code, request.MaxUses, request.RewardId)

	if err := h.PromoUC.CreatePromocode(promocode); err != nil {
		if errors.Is(err, entities.ErrPromocodeExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "promocode already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create promocode: " + err.Error()})
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

	promocode, err := h.PromoUC.ApplyPromocode(code)
	if err != nil {
		if errors.Is(err, entities.ErrPromocodeNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Promocode not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := DTOs.PromocodeResponse{
		Code:        promocode.Code,
		CurrentUses: promocode.UsesCount,
		MaxUses:     promocode.MaxUses,
	}

	c.JSON(http.StatusOK, response)
}

// GetPlayers возвращает всех игроков из бд.
func (h *Handlers) GetPlayers(c *gin.Context) {
	players, err := h.PlayerUC.GetPlayers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve players: " + err.Error()})
		return
	}

	var playersDTO []DTOs.PlayerResponse
	for _, player := range players {
		playersDTO = append(playersDTO, DTOs.PlayerResponse{
			ID:        player.ID,
			Username:  player.Username,
			Email:     player.Email,
			CreatedAt: player.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, DTOs.PlayersResponse{Players: playersDTO})
}

// GetRewards возвращает все reward из бд.
func (h *Handlers) GetRewards(c *gin.Context) {
	rewards, err := h.RewardUC.GetRewards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve rewards: " + err.Error()})
		return
	}

	var rewardsDTO []DTOs.RewardResponse
	for _, reward := range rewards {
		rewardsDTO = append(rewardsDTO, DTOs.RewardResponse{
			ID:        reward.ID,
			Reward:    reward.Reward,
			CreatedAt: reward.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, DTOs.RewardsResponse{Rewards: rewardsDTO})
}
