package dto

type GetRecommendationResponseDTO struct {
	UserUuid string `json:"user_uuid"`
	Recommendations []RecommendationDTO `json:"recommendations"`
}
