package entities

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrPromocodeNotFound       = errors.New("promocode not found")
	ErrPromocodeMaxUsesReached = errors.New("promocode has reached max uses")
	ErrPromocodeExists         = errors.New("promocode already exists")
)

type Promocode struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	MaxUses   int       `json:"max_uses" db:"max_uses"`
	UsesCount int       `json:"uses_count" db:"uses_count"`
	RewardId  uuid.UUID `json:"reward_id" db:"reward_id"`
}

func NewPromocode(code string, maxUses int, rewardId uuid.UUID) *Promocode {
	return &Promocode{
		ID:        uuid.New(),
		Code:      code,
		MaxUses:   maxUses,
		UsesCount: 0,
		RewardId:  rewardId,
	}
}
