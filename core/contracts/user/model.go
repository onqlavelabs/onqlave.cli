package user

import (
	"fmt"
	"strings"
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
)

type Users struct {
	Users []User `json:"users"`
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

type ListDetailWithAccessControl struct {
	Users []DetailWithAccessControl
	ACL   acl.ACL
}

type DetailWithAccessControl struct {
	User `json:"user"`
	ACL  acl.ACL `json:"acl"`
}

type ModelWrapper struct {
	Roles     []Role    `json:"roles"`
	Countries []Country `json:"countries"`
}

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Country struct {
	CountryCode string `json:"country_code,omitempty" validate:"required,max=20"`
	CountryName string `json:"country_name" validate:"required,max=50"`
	Icon        string `json:"icon" validate:"required,max=50"`
}

type Statistic struct {
	Total   int `json:"total"`
	Active  int `json:"active"`
	Pending int `json:"pending"`
}

type GeneralInformation struct {
	Id     string `json:"id"`
	Avatar string `json:"avatar"`
}

type TimeZone struct {
	TimeZoneCode string `json:"time_zone_code,omitempty" validate:"required,max=20"`
	TimeZoneName string `json:"time_zone_name" validate:"required,max=50"`
}

type UpdateInformation struct {
	UserID string `json:"user_id"  validate:"required"`
	Information
}

type Information struct {
	FirstName    string   `json:"first_name" `
	LastName     string   `json:"last_name" `
	EmailAddress string   `json:"email_address"   validate:"required"`
	Roles        []string `json:"role"   validate:"required"`
	CountryCode  string   `json:"country_code"   validate:"required"`
	Enable       *bool    `json:"enable"   validate:"required"`
}

func (u Information) FullName() string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", u.FirstName, u.LastName))
}

type Profile struct {
	FirstName    string `json:"first_name,omitempty"  validate:"required"`
	LastName     string `json:"last_name,omitempty"  validate:"required"`
	EmailAddress string `json:"email_address,omitempty"`
	Phone        string `json:"phone_number,omitempty"  validate:"required"`
	CountryCode  string `json:"country_code,omitempty"  validate:"required"`
	TimeZoneCode string `json:"time_zone_code,omitempty" `
}

type ProfileDetail struct {
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

type VerifyToken struct {
	Token string `json:"token" validate:"required"`
}
