package api_key

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type APIKeyResponse struct {
	common.BaseErrorResponse
	Data APIKey `json:"data"`
}

type APIKeysResponse struct {
	common.BaseErrorResponse
	Data APIKeys `json:"data"`
}

type APIKeySensitiveDataResponse struct {
	common.BaseErrorResponse
	Data APIKeySensitiveData `json:"data"`
}
