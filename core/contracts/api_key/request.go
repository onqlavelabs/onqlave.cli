package api_key

type CreateAPIKeyRequest struct {
	APIKey CreateAPIKey `json:"api_key" validate:"required"`
}

type CreateAPIKey struct {
	ApplicationID         string `json:"application_id" validate:"required"`
	ClusterID             string `json:"cluster_id" validate:"required"`
	ApplicationTechnology string `json:"application_technology" validate:"required"`
}
