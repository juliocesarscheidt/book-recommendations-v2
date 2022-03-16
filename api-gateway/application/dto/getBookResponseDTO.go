package dto

type GetBookResponseDTO struct {
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Image string `json:"image"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
