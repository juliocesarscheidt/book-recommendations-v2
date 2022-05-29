package dto

type UpdateBookWithImageRequestDTO struct {
	Uuid string `json:"uuid"`
	Title string `json:"title"`
	Author string `json:"author"`
	Genre string `json:"genre"`
	Image string `json:"image"`
}
