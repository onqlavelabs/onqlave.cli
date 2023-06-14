package contracts

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type UpdateCluster struct {
	ID            common.ArxId `json:"id" validate:"required"`
	Name          string       `json:"name" validate:"required"`
	SpendLimit    *uint64      `json:"spend_limit" validate:"required"`
	Regions       []string     `json:"regions" validate:"required,max=5"`
	RotationCycle string       `json:"rotation_cycle" validate:"required,max=50"`
	Owner         string       `json:"owner" validate:"required,max=150"`
	IsDefault     *bool        `json:"is_default" validate:"required"`
}
