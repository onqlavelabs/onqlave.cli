package auth

type RegistrationStatusDetails struct {
	Code       string `json:"status"`
	Message    string `json:"message"`
	TenantName string `json:"tenant_name"`
	TenantID   string `json:"tenant_id"`
	Token      string `json:"token"`
}

type RegistrationID struct {
	Id string `json:"id" validate:"required"`
}

type RegistrationToken struct {
	Token string `json:"token" validate:"required"`
}

type LoginDetails struct {
	UserEmail  string `json:"user_email" validate:"required"`
	TenantName string `json:"tenant_name" validate:"required"`
}

type RegistrationDetails struct {
	UserEmail    string `json:"user_email" validate:"required"`
	UserFullName string `json:"full_name" validate:"required"`
	TenantName   string `json:"tenant_name" validate:"required"`
}
