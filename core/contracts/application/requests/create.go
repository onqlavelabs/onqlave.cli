package requests

import "github.com/onqlavelabs/onqlave.cli/core/contracts/application"

type AddApplicationRequest struct {
	Application contracts.NewApplication `json:"application" validate:"required"`
}
