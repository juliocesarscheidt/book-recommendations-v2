package dto

type UpdateBookRequestDTO struct {
	Uuid string `json:"uuid,omitempty"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Genre string `json:"genre,omitempty"`
	Image string `json:"image,omitempty"`
}
