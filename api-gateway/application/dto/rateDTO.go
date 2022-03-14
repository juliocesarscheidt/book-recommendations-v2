package dto

type RateDTO struct {
	UserUuid string `json:"user_uuid,omitempty"`
	Rate float64 `json:"rate,omitempty"`
}
