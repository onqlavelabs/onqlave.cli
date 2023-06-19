package contracts

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ExistingApplicationWithDetails struct {
	ACL         acl.ACL              `json:"acl"`
	ID          common.ApplicationId `json:"application_id" validate:"required"`
	Name        string               `json:"name" validate:"required,max=150"`
	Description string               `json:"description" validate:"required,max=500"`
	Technology  string               `json:"technology" validate:"required,max=20"`
	Owner       string               `json:"owner" validate:"required,max=150"`
	APIKeys     int                  `json:"api_keys"`
	Cors        []ApplicationCors    `json:"cors" validate:"max=10"`
	Status      string               `json:"status" validate:"required"`
}
