package dto

type RecommendationDTO struct {
	BookUuid string `json:"book_uuid"`
	Rate float64 `json:"rate"`
}
