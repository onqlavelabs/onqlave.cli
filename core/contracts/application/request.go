package application

type Request struct {
	Application Application `json:"application" validate:"required"`
}
