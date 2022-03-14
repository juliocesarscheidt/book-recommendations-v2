package dto

type GetRecommendationResponseDTO struct {
	UserUuid string `json:"user_uuid,omitempty"`
	Recommendations []RecommendationDTO `json:"recommendations,omitempty"`
	// CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt string `json:"updated_at,omitempty"`
}
