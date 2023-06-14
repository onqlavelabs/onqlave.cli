package contracts

type LoginDetails struct {
	UserEmail  string `json:"user_email" validate:"required"`
	TenantName string `json:"tenant_name" validate:"required"`
}
