package DTOs

import "t_astrum/internal/promo/entities"

type PlayersResponse struct {
	Players []entities.Player `json:"players"`
}
