package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type SealApplicationResponse struct {
	common.BaseErrorResponse
	Data contracts.ApplicationStatus `json:"data"`
}
