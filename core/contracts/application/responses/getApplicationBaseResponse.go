package responses

import "github.com/onqlavelabs/onqlave.cli/core/contracts/application"

type GetApplicationBaseResponse struct {
	Data contracts.ApplicationModelWrapper `json:"data"`
}
