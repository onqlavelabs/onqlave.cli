package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type AddApplicationRequest struct {
	Application contracts.NewApplication `json:"application" validate:"required"`
}
