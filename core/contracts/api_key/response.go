package api_key

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type DetailResponse struct {
	common.BaseErrorResponse
	Data APIKey `json:"data"`
}

type ListResponse struct {
	common.BaseErrorResponse
	Data APIKeys `json:"data"`
}

type SensitiveDataResponse struct {
	common.BaseErrorResponse
	Data SensitiveData `json:"data"`
}
