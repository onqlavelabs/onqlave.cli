package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/auth"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type RegistrationResponse struct {
	common.BaseErrorResponse
	RegistrationID contracts.RegistrationID `json:"data"`
}
