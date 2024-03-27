package domain

type (
	Stock struct {
		Code         string  `json:"code"`
		Name         string  `json:"name"`
		CurrentPrice float64 `json:"current_price"`
	}
)
