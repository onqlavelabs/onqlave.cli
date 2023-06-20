package user

import "github.com/onqlavelabs/onqlave.cli/core/contracts/common"

type UpdateRequest struct {
	User UpdateInformation `json:"user"`
}

type AddRequest struct {
	User Information `json:"user"`
}

type UpdateProfileRequest struct {
	Profile ProfileDetail `json:"profile" validate:"required"`
}

type UpdateNotificationSettingRequest struct {
	NotificationSetting NotificationSetting `json:"setting"`
}

type SentVerifyEmailRequest struct {
	EmailAddress common.UserEmail `json:"email_address" validate:"required"`
}

type ResetPasswordRequest struct {
	EmailAddress common.UserEmail `json:"email_address" validate:"required"`
}

type VerifyRequest struct {
	Request VerifyToken `json:"request" validate:"required"`
}

type InviteRequest struct {
	UserID string `json:"user_id" validate:"required"`
}
