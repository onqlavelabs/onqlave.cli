package contracts

type RegistrationToken struct {
	Token string `json:"token" validate:"required"`
}
