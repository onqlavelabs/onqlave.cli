package contracts

type RegistrationDetails struct {
	UserEmail    string `json:"user_email" validate:"required"`
	UserFullName string `json:"full_name" validate:"required"`
	TenantName   string `json:"tenant_name" validate:"required"`
}
