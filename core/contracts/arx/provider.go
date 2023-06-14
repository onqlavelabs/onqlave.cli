package contracts

type ArxProvider struct {
	ID          string      `json:"id" validate:"required"`
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"required"`
	IsDefault   *bool       `json:"is_default,omitempty" validate:"required"`
	Enable      *bool       `json:"enable,omitempty" validate:"required"`
	Order       *uint8      `json:"order,omitempty" validate:"required"`
	Image       string      `json:"image" validate:"required"`
	Regions     []ArxRegion `json:"regions,omitempty" validate:"required"`
}
