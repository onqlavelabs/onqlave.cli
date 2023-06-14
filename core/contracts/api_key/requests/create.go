package requests

import "github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"

type AddAPIKeyRequest struct {
	APIKey contracts.NewAPIKey `json:"api_key" validate:"required"`
}
