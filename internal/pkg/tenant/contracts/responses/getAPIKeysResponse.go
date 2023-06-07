package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetAPIKeysResponse struct {
	common.BaseErrorResponse
	Data contracts.GetAPIKeysResponseWrapper `json:"data"`
}

type GetAPIKeyBaseInformationResponse struct {
	common.BaseErrorResponse
	Data contracts.GetAPIKeyBaseResponse `json:"data"`
}
