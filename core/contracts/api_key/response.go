package api_key

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type APIKeyDetailResponse struct {
	common.BaseErrorResponse
	Data APIKeyDetail `json:"data"`
}

type GetAPIKeysResponse struct {
	common.BaseErrorResponse
	Data GetAPIKeysResponseWrapper `json:"data"`
}

type APIKeySensitiveDataResponse struct {
	common.BaseErrorResponse
	Data SensitiveData `json:"data"`
}
