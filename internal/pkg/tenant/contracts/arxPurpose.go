package contracts

type ArxPurpose struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	IsDefault *bool  `json:"is_default,omitempty" validate:"required"`
	Enable    *bool  `json:"enable,omitempty" validate:"required"`
	Order     *uint8 `json:"order,omitempty" validate:"required"`
}
