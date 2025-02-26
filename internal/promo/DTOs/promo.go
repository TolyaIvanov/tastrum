package DTOs

type CreatePromocodeRequest struct {
	Code    string `json:"code"`
	MaxUses int    `json:"max_uses"`
}

type PromocodeResponse struct {
	Code    string `json:"code"`
	MaxUses int    `json:"max_uses"`
}
