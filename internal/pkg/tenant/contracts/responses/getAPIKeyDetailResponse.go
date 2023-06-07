package responses

import (
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"
)

type GetAPIKeyDetailResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKeyDetail `json:"data"`
}

type GetAPIKeySensitiveInfoResponse struct {
	common.BaseErrorResponse
	Data contracts.APIKeySensitive `json:"data"`
}
