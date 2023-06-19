package application

type ApplicationRequest struct {
	Application Application `json:"application" validate:"required"`
}
