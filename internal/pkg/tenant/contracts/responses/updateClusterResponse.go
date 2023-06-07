package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type UpdateClusterResponse struct {
	common.BaseErrorResponse
	Data contracts.ExistingClusterWithDetails `json:"data"`
}