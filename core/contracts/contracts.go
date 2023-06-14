package contracts

import (
	"fmt"
	"strings"
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type AddTenantRequest struct {
	Tenant Tenant `json:"tenant" validate:"required"`
}

type AddTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type TenantId struct {
	Id string `json:"id" validate:"required"`
}

type Tenant struct {
	Id          common.TenantId `json:"tenant_id" validate:"required"`
	Name        string          `json:"name" validate:"required,min=4,max=100"`
	Description string          `json:"description" validate:"max=500"`
	Disable     bool            `json:"disable"`
	OwnerEmail  string          `json:"owner_email" validate:"email,required"`
	RequestId   string          `json:"request_id" validate:"required"`
}

type TenantInfo struct {
	Id         common.TenantId `json:"tenant_id,omitempty"`
	Name       string          `json:"tenant_name" validate:"required,min=4,max=100"`
	Label      string          `json:"tenant_label"  validate:"required"`
	OwnerEmail string          `json:"owner_email,omitempty"`
	CreatedOn  time.Time       `json:"created_on,omitempty"`
	ACL        contracts.ACL   `json:"acl,omitempty"`
}

type TenantDashboard struct {
	Insight DasboardInsight  `json:"insight"`
	Events  []DashboardEvent `json:"events"`
}

type DasboardInsight struct {
	NumberOfArx          int `json:"number_of_arx"`
	NumberOfApplications int `json:"number_of_applications"`
	NumberOfApiKeys      int `json:"number_of_apikeys"`
	NumberOfUsers        int `json:"number_of_users"`
}

type DashboardEvent struct {
	Operation string    `json:"operations"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	IsError   bool      `json:"is_error"`
}

type UpdateTenantRequest struct {
	Tenant TenantInfo `json:"tenant" validate:"required"`
}

type UpdateTenantResponse struct {
	common.BaseErrorResponse
	Tenant TenantInfo `json:"data"`
}

type DisableTenantRequest struct {
	Disable *bool `json:"disable" validate:"required"`
}

type DisableTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type DeleteTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type GetTenantResponse struct {
	common.BaseErrorResponse
	Tenant TenantInfo `json:"data"`
}

type GetTenantDashboardResponse struct {
	common.BaseErrorResponse
	Dashboard TenantDashboard `json:"data"`
}

type GetTenantsResponse struct {
	common.BaseErrorResponse
	Tenants []Tenant `json:"data"`
}

type ApplicationClientType struct {
	Id    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}

type GetProfileResponse struct {
	common.BaseErrorResponse
	Profile Profile `json:"data"`
}

type GetUserNotificationSettingsResponse struct {
	common.BaseErrorResponse
	Settings []NotificationSetting `json:"data" validate:"required"`
}

type UpdateProfileRequest struct {
	Profile UserProfile `json:"profile" validate:"required"`
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

type GetSecurityEventResponse struct {
	common.BaseErrorResponse
	SercurityEvents []SecurityEvent `json:"data"`
}

type UpdateProfileResponse struct {
	common.BaseErrorResponse
	Data UserProfile `json:"data"`
}

type UpdateNotificationSettingResponse struct {
	common.BaseErrorResponse
	UserId common.UserId `json:"data"`
}

type AddUserRequest struct {
	User UserInformation `json:"user"`
}

type AddUserResponse struct {
	common.BaseErrorResponse
	Data UserWithAccessControl `json:"data"`
}

type UpdateUserRequest struct {
	User UserUpdateInformation `json:"user"`
}

type UpdateUserResponse struct {
	common.BaseErrorResponse
	Data UserWithAccessControl `json:"data"`
}

type UserGeneralInformation struct {
	Id     string `json:"id"`
	Avatar string `json:"avatar"`
}

type GetUsersResponse struct {
	Users         []UserWithAccessControl `json:"users"`
	Model         UserModelWrapper        `json:"model"`
	ACL           contracts.ACL           `json:"acl"`
	UserStatistic UserStatistic           `json:"statistics"`
}

type UserModelWrapper struct {
	Roles     []Role    `json:"roles"`
	Countries []Country `json:"countries"`
}

// type APIKey struct {
// 	Id                string        `json:"id" validate:"required"`
// 	Key               string        `json:"api_key" validate:"required"`
// 	SignKey           string        `json:"sign_key" validate:"required"`
// 	CryptoKey         string        `json:"crypto_key" validate:"required"`
// 	ClientKey         string        `json:"client_key" validate:"required"`
// 	Name              string        `json:"name" validate:"required"`
// 	Description       string        `json:"description" validate:"required"`
// 	Rotated           *bool         `json:"rotated" validate:"required"`
// 	ApplicationId     ApplicationId `json:"application_id" validate:"required"`
// 	ActiveVersion     APIKeyVersion `json:"active_version" validate:"required"`
// 	LastActiveVersion APIKeyVersion `json:"last_active_version" validate:"required"`
// }

type InviteUserRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

type APIKeyVersion struct {
	Number    string    `json:"version_number" validate:"required"`
	CreatedAt time.Time `json:"created_at" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
}

type UsersWithAccessControl struct {
	Users []UserWithAccessControl
	ACL   contracts.ACL
}

type UserWithAccessControl struct {
	User `json:"user"`
	ACL  contracts.ACL `json:"acl"`
}

type UserStatistic struct {
	Total   int `json:"total"`
	Active  int `json:"active"`
	Pending int `json:"pending"`
}

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type User struct {
	ID           string    `json:"id" validate:"required"`
	FullName     string    `json:"full_name"`
	Avatar       string    `json:"avatar"`
	Status       string    `json:"status,omitempty"`
	EmailAddress string    `json:"email_address" validate:"required"`
	TenantID     string    `json:"tenant_id" `
	Role         []string  `json:"role"`
	Disable      bool      `json:"disable"`
	Country      string    `json:"country_code,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Country struct {
	CountryCode string `json:"country_code,omitempty" validate:"required,max=20"`
	CountryName string `json:"country_name" validate:"required,max=50"`
	Icon        string `json:"icon" validate:"required,max=50"`
}

type TimeZone struct {
	TimeZoneCode string `json:"time_zone_code,omitempty" validate:"required,max=20"`
	TimeZoneName string `json:"time_zone_name" validate:"required,max=50"`
}

type UserUpdateInformation struct {
	UserID string `json:"user_id"  validate:"required"`
	UserInformation
}

type UserInformation struct {
	FirstName    string   `json:"first_name" `
	LastName     string   `json:"last_name" `
	EmailAddress string   `json:"email_address"   validate:"required"`
	Roles        []string `json:"role"   validate:"required"`
	CountryCode  string   `json:"country_code"   validate:"required"`
	Enable       *bool    `json:"enable"   validate:"required"`
}

func (u UserInformation) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", u.FirstName, u.LastName))
}

type UserProfile struct {
	FirstName    string `json:"first_name,omitempty"  validate:"required"`
	LastName     string `json:"last_name,omitempty"  validate:"required"`
	EmailAddress string `json:"email_address,omitempty"`
	Phone        string `json:"phone_number,omitempty"  validate:"required"`
	CountryCode  string `json:"country_code,omitempty"  validate:"required"`
	TimeZoneCode string `json:"time_zone_code,omitempty" `
}

type Profile struct {
	FirstName    string     `json:"first_name,omitempty"`
	LastName     string     `json:"last_name,omitempty"`
	EmailAddress string     `json:"email_address,omitempty" validate:"required"`
	Phone        string     `json:"phone_number,omitempty"`
	CountryCode  string     `json:"country_code,omitempty"`
	TimeZoneCode string     `json:"time_zone_code,omitempty"`
	Countries    []Country  `json:"countries" validate:"required"`
	TimeZones    []TimeZone `json:"time_zones" validate:"required"`
}

type NotificationSetting struct {
	Code        string `json:"code" validate:"required,max=50"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Enable      *bool  `json:"enable" validate:"required"`
}

type ResetPasswordInfo struct {
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	ResetLink    string `json:"reset_link"`
}

type VerifyEmailInfo struct {
	FullName     string `json:"full_name"`
	EmailAddress string `json:"email_address"`
	VerifyLink   string `json:"verify_link"`
}

type InvitationEmailInfo struct {
	FullName       string `json:"full_name"`
	EmailAddress   string `json:"email_address"`
	InvitationLink string `json:"invitation_link"`
}

type SecurityEvent struct {
	Action       string    `json:"action"`
	IpAddress    string    `json:"ip_address"`
	Location     string    `json:"location"`
	ActivityTime time.Time `json:"activity_time"`
}

type UserVerifyRequest struct {
	Request VerifyToken `json:"request" validate:"required"`
}

type VerifyToken struct {
	Token string `json:"token" validate:"required"`
}
