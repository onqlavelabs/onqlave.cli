package application

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ArchiveApplicationResponse struct {
	common.BaseErrorResponse
	Data ApplicationStatus `json:"data"`
}

type AddApplicationResponse struct {
	common.BaseErrorResponse
	Data ExistingApplicationWithDetails `json:"data"`
}

type DisableApplicationResponse struct {
	common.BaseErrorResponse
	Data ApplicationStatus `json:"data"`
}

type GetApplicationResponse struct {
	common.BaseErrorResponse
	Data ExistingApplicationWithDetails `json:"data"`
}

type GetApplicationsResponse struct {
	common.BaseErrorResponse
	Data GetApplications `json:"data"`
}

type GetApplicationBaseResponse struct {
	Data ApplicationModelWrapper `json:"data"`
}

type UpdateApplicationResponse struct {
	common.BaseErrorResponse
	Data ExistingApplicationWithDetails `json:"data"`
}
