package api_key

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type AddAPIKeyResponse struct {
	common.BaseErrorResponse
	Data APIKey `json:"data"`
}

type GetAPIKeysResponse struct {
	common.BaseErrorResponse
	Data GetAPIKeysResponseWrapper `json:"data"`
}

type GetAPIKeyBaseInformationResponse struct {
	common.BaseErrorResponse
	Data GetAPIKeyBaseResponse `json:"data"`
}
type GetAPIKeyDetailResponse struct {
	common.BaseErrorResponse
	Data APIKeyDetail `json:"data"`
}

type GetAPIKeySensitiveInfoResponse struct {
	common.BaseErrorResponse
	Data SensitiveData `json:"data"`
}
