package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/auth"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type RegistrationStatusResponse struct {
	common.BaseErrorResponse
	Status contracts.RegistrationStatusDetails `json:"data"`
}
