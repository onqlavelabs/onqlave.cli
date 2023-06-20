package user

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type GetListResponse struct {
	common.BaseErrorResponse
	Data ListResponse `json:"data"`
}

type ListResponse struct {
	Users         []DetailWithAccessControl `json:"users"`
	Model         ModelWrapper              `json:"model"`
	ACL           acl.ACL                   `json:"acl"`
	UserStatistic Statistic                 `json:"statistics"`
}

type UpdateResponse struct {
	common.BaseErrorResponse
	Data DetailWithAccessControl `json:"data"`
}

type AddUserResponse struct {
	common.BaseErrorResponse
	Data DetailWithAccessControl `json:"data"`
}

type GetSecurityEventResponse struct {
	common.BaseErrorResponse
	SecurityEvents []SecurityEvent `json:"data"`
}

type UpdateProfileResponse struct {
	common.BaseErrorResponse
	Data ProfileDetail `json:"data"`
}

type UpdateNotificationSettingResponse struct {
	common.BaseErrorResponse
	UserId common.UserId `json:"data"`
}

type GetProfileResponse struct {
	common.BaseErrorResponse
	Profile Profile `json:"data"`
}

type GetUserNotificationSettingsResponse struct {
	common.BaseErrorResponse
	Settings []NotificationSetting `json:"data" validate:"required"`
}
