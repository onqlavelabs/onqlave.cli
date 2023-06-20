package application

type Request struct {
	Application RequestApplication `json:"application" validate:"required"`
}
