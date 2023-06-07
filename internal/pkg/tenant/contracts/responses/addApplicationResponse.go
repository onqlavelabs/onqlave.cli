package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type AddApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingApplicationWithDetails `json:"data"`
}
