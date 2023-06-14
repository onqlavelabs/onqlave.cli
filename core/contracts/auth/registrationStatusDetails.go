package contracts

type RegistrationStatusDetails struct {
	Code       string `json:"status"`
	Message    string `json:"message"`
	TenantName string `json:"tenant_name"`
	TenantID   string `json:"tenant_id"`
	Token      string `json:"token"`
}
