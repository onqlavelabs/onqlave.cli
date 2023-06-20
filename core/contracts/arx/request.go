package contracts

type AddRequest struct {
	Arx NewArx `json:"cluster" validate:"required"`
}

type UpdateRequest struct {
	Arx UpdateArx `json:"cluster" validate:"required"`
}
