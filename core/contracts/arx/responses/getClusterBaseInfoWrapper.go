package responses

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/arx"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetClusterBaseInfoWrapper struct {
	common.BaseErrorResponse
	Data contracts.ArxModelWrapper `json:"data"`
}
