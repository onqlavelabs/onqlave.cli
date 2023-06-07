package contracts

type ApplicationModelWrapper struct {
	Technologies []ApplicationTechnology `json:"technologies" validate:"required"`
}
