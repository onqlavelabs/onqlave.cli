package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type LoginRequest struct {
	LoginDetails contracts.LoginDetails `json:"login" validate:"required"`
}
