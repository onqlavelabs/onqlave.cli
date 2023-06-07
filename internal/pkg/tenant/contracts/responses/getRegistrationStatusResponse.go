package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type RegistrationStatusResponse struct {
	common.BaseErrorResponse
	Status contracts.RegistrationStatusDetails `json:"data"`
}
