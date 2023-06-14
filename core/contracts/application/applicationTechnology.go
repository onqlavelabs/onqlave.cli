package contracts

type ApplicationTechnology struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Cors        bool   `json:"cors" validate:"required"`
	Order       int16  `json:"order" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Enable      bool   `json:"enable" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
}
