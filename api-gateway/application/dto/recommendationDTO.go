package dto

type RecommendationDTO struct {
	BookUuid string `json:"book_uuid,omitempty"`
	Rate float64 `json:"rate,omitempty"`
}
