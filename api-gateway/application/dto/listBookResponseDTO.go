package dto

type ListBookResponseDTO struct {
	Books []GetBookResponseDTO `json:"books"`
}
