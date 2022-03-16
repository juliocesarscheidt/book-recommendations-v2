package dto

type CreateBookRequestDTO struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Image string `json:"image"`
}
