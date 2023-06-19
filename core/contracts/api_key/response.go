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
type APIKeyDetailResponse struct {
	common.BaseErrorResponse
	Data APIKeyDetail `json:"data"`
}

type APIKeySensitiveDataResponse struct {
	common.BaseErrorResponse
	Data SensitiveData `json:"data"`
}
