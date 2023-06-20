package auth

type RegistrationStatusRequest struct {
	Request RegistrationToken `json:"request" validate:"required"`
}

type LoginRequest struct {
	LoginDetails LoginDetails `json:"login" validate:"required"`
}

type RegistrationRequest struct {
	Registration RegistrationDetails `json:"registration" validate:"required"`
}
