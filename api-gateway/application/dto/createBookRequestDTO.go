package dto

type CreateBookRequestDTO struct {
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Genre string `json:"genre,omitempty"`
	Image string `json:"image,omitempty"`
}
