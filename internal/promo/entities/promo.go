package entities

type Promocode struct {
	Code      string `json:"code"`
	MaxUses   int    `json:"max_uses"`
	UsesCount int    `json:"uses_count"`
	Reward    string `json:"reward"`
}
