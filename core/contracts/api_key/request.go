package api_key

type CreateAPIKeyRequest struct {
	APIKey CreateAPIKey `json:"api_key" validate:"required"`
}
