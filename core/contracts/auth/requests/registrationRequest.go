package requests

import "github.com/onqlavelabs/onqlave.cli/core/contracts/auth"

type RegistrationRequest struct {
	Registration contracts.RegistrationDetails `json:"registration" validate:"required"`
}
