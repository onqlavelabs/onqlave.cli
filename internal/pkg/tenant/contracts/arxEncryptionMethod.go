package contracts

type ArxEncryptionMethod struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDefault   *bool  `json:"is_default,omitempty" validate:"required"`
	Enable      *bool  `json:"enable,omitempty" validate:"required"`
	Order       *uint8 `json:"order,omitempty" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}
