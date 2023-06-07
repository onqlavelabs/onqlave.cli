package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type AddAPIKeyRequest struct {
	APIKey contracts.NewAPIKey `json:"api_key" validate:"required"`
}
