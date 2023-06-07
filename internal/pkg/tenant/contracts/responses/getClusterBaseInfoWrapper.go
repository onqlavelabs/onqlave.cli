package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetClusterBaseInfoWrapper struct {
	common.BaseErrorResponse
	Data contracts.ArxModelWrapper `json:"data"`
}
