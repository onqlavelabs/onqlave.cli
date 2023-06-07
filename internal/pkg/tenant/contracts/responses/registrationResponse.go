package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type RegistrationResponse struct {
	common.BaseErrorResponse
	RegistrationID contracts.RegistrationID `json:"data"`
}
