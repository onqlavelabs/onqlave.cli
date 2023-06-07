package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type AddAPIKeyResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKey `json:"data"`
}
