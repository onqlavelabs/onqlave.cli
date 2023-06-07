package responses

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type GetApplicationBaseResponse struct {
	Data contracts.ApplicationModelWrapper `json:"data"`
}
