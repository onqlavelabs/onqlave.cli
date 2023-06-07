package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type RegistrationStatusRequest struct {
	Request contracts.RegistrationToken `json:"request" validate:"required"`
}
