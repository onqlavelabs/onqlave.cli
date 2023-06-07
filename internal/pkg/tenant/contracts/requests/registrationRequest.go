package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type RegistrationRequest struct {
	Registration contracts.RegistrationDetails `json:"registration" validate:"required"`
}
