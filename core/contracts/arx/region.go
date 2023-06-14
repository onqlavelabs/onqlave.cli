package contracts

type ArxRegionOptimisation struct {
	Value   int32  `json:"value" validate:"required"`
	Message string `json:"message" validate:"required,max=50"`
}

type ArxRegion struct {
	ID           string                `json:"id" validate:"required"`
	Name         string                `json:"name" validate:"required"`
	IsDefault    *bool                 `json:"is_default,omitempty" validate:"required"`
	Enable       *bool                 `json:"enable,omitempty" validate:"required"`
	Order        *uint8                `json:"order,omitempty" validate:"required"`
	Icon         string                `json:"icon" validate:"required"`
	Optimisation ArxRegionOptimisation `json:"optimisation" validate:"required"`
}
