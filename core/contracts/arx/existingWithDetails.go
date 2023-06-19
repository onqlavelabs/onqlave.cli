package contracts

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ExistingClusterWithDetails struct {
	ACL                 acl.ACL      `json:"acl"`
	ID                  common.ArxId `json:"id" validate:"required"`
	Name                string       `json:"name" validate:"required"`
	SpendLimit          uint64       `json:"spend_limit"`
	Description         string       `json:"description" validate:"required,min=4,max=500"`
	Purpose             string       `json:"purpose" validate:"required,max=50"`
	Plan                string       `json:"plan" validate:"required,max=50"`
	Provider            string       `json:"provider" validate:"required,max=50"`
	Regions             []string     `json:"regions" validate:"required,max=5"`
	AvailabilityMessage string       `json:"availability_message"`
	EncryptionMethod    string       `json:"encryption_method" validate:"required,max=50"`
	RotationCycle       string       `json:"rotation_cycle" validate:"required,max=50"`
	Owner               string       `json:"owner" validate:"required,max=150"`
	IsDefault           bool         `json:"is_default"`
	Status              string       `json:"status" validate:"required,max=50"`
}
