package contracts

type RegistrationID struct {
	Id string `json:"id" validate:"required"`
}
