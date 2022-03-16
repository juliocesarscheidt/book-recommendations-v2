package dto

type UpdateBookRequestDTO struct {
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Image string `json:"image"`
}
