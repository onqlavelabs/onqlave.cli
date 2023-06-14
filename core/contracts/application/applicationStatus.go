package contracts

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ApplicationStatus struct {
	ID        common.ApplicationId `json:"application_id" validate:"required"`
	State     string               `json:"data"`
	Message   string               `json:"message"`
	Operation string               `json:"operation"`
	IsError   bool                 `json:"is_error"`
	UpdatedAt time.Time            `json:"update_time"`
}
