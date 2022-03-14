package dto

type GetBookResponseDTO struct {
	Uuid string `json:"uuid,omitempty"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Genre string `json:"genre,omitempty"`
	Image string `json:"image,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
