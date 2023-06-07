package domainevents

import (
	"github.com/google/uuid"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
)

type DomainEvent interface {
	SetEventId(id uuid.UUID)
	GetEventId() uuid.UUID
	Topic() string
	Metadata() map[string]string
}

type RegistrationRequestReceived struct {
	EventId          uuid.UUID `json:"event_id" validate:"required,max=100"`
	EmailAddress     string    `json:"email" validate:"email,required,min=4,max=50"`
	Token            string    `json:"token" validate:"required,min=4,max=100"`
	LinkValidityTime int32     `json:"validity" validate:"min=15,max=120"`
	Link             string    `json:"link" validate:"required,min=10,max=500"`
	Operation        string    `json:"operation" validate:"required,min=10,max=50"`
	TenantName       string    `json:"tenant_name" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestReceived) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *RegistrationRequestReceived) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *RegistrationRequestReceived) Topic() string {
	return "RegistrationRequestReceived"
}

func (d *RegistrationRequestReceived) Metadata() map[string]string {
	return nil
}

type RegistrationRequestExpired struct {
	EventId uuid.UUID `json:"event_id" validate:"required,max=100"`
	Token   string    `json:"token" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestExpired) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *RegistrationRequestExpired) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *RegistrationRequestExpired) Metadata() map[string]string {
	return nil
}

func (d *RegistrationRequestExpired) Topic() string {
	return "RegistrationRequestExpired"
}

type RegistrationEmailSent struct {
	EventId uuid.UUID `json:"event_id" validate:"required,max=100"`
	Token   string    `json:"token" validate:"required,min=4,max=100"`
	Message string    `json:"message" validate:"required"`
	IsError bool      `json:"is_error"`
}

type TenantAdded struct {
	EventId           uuid.UUID `json:"event_id" validate:"required,max=100"`
	Name              string    `json:"tenant_name" validate:"required,min=4,max=100"`
	Id                string    `json:"tenant_id" validate:"required,min=4,max=100"`
	Message           string    `json:"message" validate:"required"`
	IsError           bool      `json:"is_error"`
	RequestId         string    `json:"request_id" validate:"required,min=4,max=100"`
	OwnerEmailAddress string    `json:"owner_email_address" validate:"required,email"`
}

func (d *TenantAdded) Topic() string {
	return "TenantAdded"
}

func (d *TenantAdded) Metadata() map[string]string {
	return nil
}

func (d *TenantAdded) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *TenantAdded) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *RegistrationEmailSent) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *RegistrationEmailSent) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *RegistrationEmailSent) Topic() string {
	return "RegistrationEmailSent"
}

func (d *RegistrationEmailSent) Metadata() map[string]string {
	return nil
}

type RegistrationRequestCompleted struct {
	EventId      uuid.UUID `json:"event_id" validate:"required,max=100"`
	EmailAddress string    `json:"email" validate:"email,required,min=4,max=50"`
	Token        string    `json:"token" validate:"required,min=4,max=100"`
}

func (d *RegistrationRequestCompleted) Topic() string {
	return "RegistrationRequestCompleted"
}

func (d *RegistrationRequestCompleted) Metadata() map[string]string {
	return nil
}

func (d *RegistrationRequestCompleted) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *RegistrationRequestCompleted) GetEventId() uuid.UUID {
	return d.EventId
}

type ClusterStateChanged struct {
	EventId    uuid.UUID                   `json:"event_id" validate:"required,max=100"`
	Id         common.ArxId                `json:"cluster_id" validate:"required,min=10,max=250"`
	State      common.ArxProvisioningState `json:"state" validate:"required,min=4,max=20"`
	TenantId   common.TenantId             `json:"tenant_id" validate:"required,min=10,max=200"`
	Endpoint   common.ArxEndpoint          `json:"endpoint" validate:"required,min=10,max=1000"`
	Provider   common.ArxProvider          `json:"provider" validate:"required"`
	Type       common.ArxType              `json:"cluster_type" validate:"required"`
	Attributes map[string]string           `json:"attributes" validate:"max=20"`
	Message    string                      `json:"message" validate:"max=500"`
	IsError    bool                        `json:"is_error"`
}

func (d *ClusterStateChanged) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ClusterStateChanged) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ClusterStateChanged) Topic() string {
	return "ClusterStateChanged"
}

func (d *ClusterStateChanged) Metadata() map[string]string {
	return nil
}

type ClusterPayload struct {
	TenantId common.TenantId    `json:"tenant_id" validate:"required,max=150"`
	Id       common.ArxId       `json:"cluster_id" validate:"required,max=150"`
	Name     string             `json:"cluster_name" validate:"required,max=150"`
	Provider common.ArxProvider `json:"cluster_provider" validate:"required,max=10"`
	Type     common.ArxType     `json:"cluster_type" validate:"oneof=1 2 3"`
	Purpose  common.ArxPurpose  `json:"cluster_purpose" validate:"oneof=1 2 3"`
	Region   common.ArxRegion   `json:"cluster_region" validate:"oneof=1 2"`
	UserId   string             `json:"user_id" validate:"required"`
}

type ClusterAdded struct {
	EventId uuid.UUID `json:"event_id" validate:"required,max=100"`
	ClusterPayload
}

func (d *ClusterAdded) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ClusterAdded) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ClusterAdded) Topic() string {
	return "ClusterAdded"
}

func (d *ClusterAdded) Metadata() map[string]string {
	return nil
}

type ClusterDeleted struct {
	EventId uuid.UUID `json:"event_id" validate:"required,max=100"`
	ClusterPayload
}

func (d *ClusterDeleted) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ClusterDeleted) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ClusterDeleted) Topic() string {
	return "ClusterDeleted"
}

func (d *ClusterDeleted) Metadata() map[string]string {
	return nil
}

type EmailVerificationRequestReceived struct {
	EventId          uuid.UUID `json:"event_id" validate:"required,max=100"`
	EmailAddress     string    `json:"email_address" validate:"email,required,min=4,max=50"`
	VerificationLink string    `json:"verification_link" validate:"required,min=10,max=200"`
	FullNameMasked   string    `json:"full_name_masked" validate:"required,min=10,max=200"`
}

func (d *EmailVerificationRequestReceived) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *EmailVerificationRequestReceived) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *EmailVerificationRequestReceived) Topic() string {
	return "EmailVerificationRequestReceived"
}

func (d *EmailVerificationRequestReceived) Metadata() map[string]string {
	return nil
}

type ResetPasswordRequestReceived struct {
	EventId        uuid.UUID `json:"event_id" validate:"required,max=100"`
	EmailAddress   string    `json:"email_address" validate:"email,required,min=4,max=50"`
	ResetLink      string    `json:"verification_link" validate:"required,min=10,max=200"`
	FullNameMasked string    `json:"full_name_masked" validate:"required,min=10,max=200"`
}

func (d *ResetPasswordRequestReceived) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ResetPasswordRequestReceived) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ResetPasswordRequestReceived) Topic() string {
	return "EmailVerificationRequestReceived"
}

func (d *ResetPasswordRequestReceived) Metadata() map[string]string {
	return nil
}

type TenantUpdated struct {
	EventId     uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantId    string    `json:"tenant_id" validate:"required,min=10,max=200"`
	TenantName  string    `json:"tenant_name" validate:"required,min=10,max=200"`
	TenantLabel string    `json:"tenant_label" validate:"required,min=10,max=200"`
}

func (d *TenantUpdated) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *TenantUpdated) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *TenantUpdated) Topic() string {
	return "TenantUpdated"
}

func (d *TenantUpdated) Metadata() map[string]string {
	return nil
}

type TenantDeactivated struct {
	EventId           uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID          string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	OwnerEmailAddress string    `json:"owner_email_address" validate:"email,required,min=4,max=50"`
	UserID            string    `json:"user_id" validate:"required,max=150"`
}

func (d *TenantDeactivated) Topic() string {
	return "TenantDeactivated"
}

func (d *TenantDeactivated) Metadata() map[string]string {
	return nil
}

type TenantReactivated struct {
	EventId          uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID         string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	SubscriptionType string    `json:"subscription_type" validate:"required,max=100"`
	UserID           string    `json:"user_id" validate:"required,max=150"`
}

func (d *TenantReactivated) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *TenantReactivated) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *TenantReactivated) Topic() string {
	return "TenantReactivated"
}

func (d *TenantReactivated) Metadata() map[string]string {
	return nil
}

func (d *TenantDeactivated) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *TenantDeactivated) GetEventId() uuid.UUID {
	return d.EventId
}

type UserAdded struct {
	EventId        uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID       string    `json:"tenant_id" validate:"required,max=150"`
	AddedUserID    string    `json:"added_user_id" validate:"required,max=150"`
	InvitationLink string    `json:"invitation_link" validate:"required,max=150"`
	FullNameMasked string    `json:"full_name_masked" validate:"required,max=150"`
	EmailAddress   string    `json:"email_address" validate:"required,max=150"`
	UserID         string    `json:"user_id" validate:"required,max=150"`
}

func (d *UserAdded) Topic() string {
	return "UserAdded"
}

func (d *UserAdded) Metadata() map[string]string {
	return nil
}

type UserInvitationReceived struct {
	EventId        uuid.UUID `json:"event_id" validate:"required,max=100"`
	EmailAddress   string    `json:"email_address" validate:"required,max=150"`
	InvitationLink string    `json:"invitation_link" validate:"required,max=150"`
	FullNameMasked string    `json:"full_name_masked" validate:"required,max=150"`
}

func (d *UserInvitationReceived) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *UserInvitationReceived) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *UserInvitationReceived) Topic() string {
	return "UserInvitationReceived"
}

func (d *UserInvitationReceived) Metadata() map[string]string {
	return nil
}

type UserDeleted struct {
	EventId        uuid.UUID `json:"event_id" validate:"required,max=100"`
	UserID         string    `json:"user_id" validate:"required"`
	TenantID       string    `json:"tenant_id" validate:"required"`
	EmailAddress   string    `json:"email_address" validate:"email,required,max=150"`
	FullNameMasked string    `json:"full_name_masked" validate:"required,max=150"`
}

func (d *UserDeleted) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *UserDeleted) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *UserDeleted) Topic() string {
	return "UserDeleted"
}

func (d *UserDeleted) Metadata() map[string]string {
	return nil
}

func (d *UserAdded) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *UserAdded) GetEventId() uuid.UUID {
	return d.EventId
}

type BillingNotificationSettingsUpdated struct {
	EventId                uuid.UUID `json:"event_id" validate:"required,max=100"`
	NotificationsSettingId string    `json:"notifications_setting_id" validate:"required,max=250"`
	UserID                 string    `json:"user_id" validate:"required"`
	SubSystem              string    `json:"sub_system" validate:"max=100"`
	TenantID               string    `json:"tenant_id" validate:"required"`
}

func (d *BillingNotificationSettingsUpdated) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *BillingNotificationSettingsUpdated) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *BillingNotificationSettingsUpdated) Topic() string {
	return "BillingNotificationSettingsUpdated"
}

func (d *BillingNotificationSettingsUpdated) Metadata() map[string]string {
	return nil
}

type ClustersSealed struct {
	EventId           uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID          string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	OwnerEmailAddress string    `json:"owner_email_address" validate:"email,required,min=4,max=50"`
	UserID            string    `json:"user_id" validate:"required,max=150"`
}

func (d *ClustersSealed) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ClustersSealed) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ClustersSealed) Topic() string {
	return "ClustersDeleted"
}

func (d *ClustersSealed) Metadata() map[string]string {
	return nil
}

type FeedbackAdded struct {
	EventId   uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID  string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	UserID    string    `json:"user_id" validate:"required,max=150"`
	FromEmail string    `json:"from_email" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Page      string    `json:"page" validate:"required"`
	Like      *bool     `json:"like" validate:"required"`
}

func (d *FeedbackAdded) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *FeedbackAdded) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *FeedbackAdded) Topic() string {
	return "FeedbackAdded"
}

func (d *FeedbackAdded) Metadata() map[string]string {
	return nil
}
