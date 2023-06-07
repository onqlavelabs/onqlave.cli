package requests

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type UpdateApplicationRequest struct {
	Application contracts.UpdateApplication `json:"application" validate:"required"`
}
