package api_key

type CreateRequest struct {
	APIKey CreateAPIKey `json:"api_key" validate:"required"`
}
