package application

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type DetailResponse struct {
	common.BaseErrorResponse
	Data Application `json:"data"`
}

type BaseResponse struct {
	Data Technologies `json:"data"`
}

type ListResponse struct {
	common.BaseErrorResponse
	Data Applications `json:"data"`
}

type StatusResponse struct {
	common.BaseErrorResponse
	Data Status `json:"data"`
}
