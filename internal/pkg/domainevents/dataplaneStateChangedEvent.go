package domainevents

import (
	"github.com/google/uuid"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/common"
)

type EventType string

const (
	ProvisionClusterChanged       EventType = "provisionclusterstatechanged"
	SealClusterChanged            EventType = "sealclusterstatechanged"
	UnsealClusterChanged          EventType = "unsealclusterstatechanged"
	DeleteClusterChanged          EventType = "deleteclusterstatechanged"
	RotateClusterKeysChanged      EventType = "rotateclusterkeysstatechanged"
	UpdateClusterChanged          EventType = "updateclusterstatechanged"
	CreateApplicationChanged      EventType = "createapplicationstatechanged"
	DisableApplicationChanged     EventType = "disableapplicationstatechanged"
	EnableApplicationChanged      EventType = "enableapplicationstatechanged"
	ArchiveApplicationChanged     EventType = "archiveapplicationstatechanged"
	UpgradeClusterWorkloadChanged EventType = "upgradeclusterworkloadstatechanged"
	CreateAPIKeyChanged           EventType = "createapikeystatechanged"
	DeleteAPIKeyChanged           EventType = "deleteapikeystatechanged"
	RotateAPIKeyChanged           EventType = "rotateapikeystatechanged"
)

type DataplaneStateChanged struct {
	EventId    uuid.UUID         `json:"event_id" validate:"required,max=100"`
	EventType  EventType         `json:"event_type" validate:"max=50"`
	Event      []byte            `json:"command" validate:"max=2000"`
	Id         string            `json:"id" validate:"required,min=10,max=250"`
	TenantId   common.TenantId   `json:"tenant_id" validate:"required,min=10,max=250"`
	Attributes map[string]string `json:"attributes" validate:"max=20"`
	TraceId    string            `json:"trace_id" validate:"required,min=10,max=150"`
}

func (d *DataplaneStateChanged) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *DataplaneStateChanged) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *DataplaneStateChanged) Topic() string {
	return "DataplaneStateChanged"
}

func (d *DataplaneStateChanged) Metadata() map[string]string {
	return d.Attributes
}

type ProvisionClusterStateChanged struct {
	State    string             `json:"state" validate:"required,min=4,max=20"`
	Region   string             `json:"region" validate:"required,max=50"`
	Endpoint common.ArxEndpoint `json:"endpoint" validate:"required,min=10,max=500"`
	Message  string             `json:"message" validate:"max=500"`
	IsError  bool               `json:"is_error"`
}

type DeleteClusterStateChanged struct {
	ArxId   common.ArxId `json:"cluster_id"`
	State   string       `json:"state" validate:"required,min=4,max=20"`
	Region  string       `json:"region" validate:"required,max=50"`
	Message string       `json:"message" validate:"max=500"`
	IsError bool         `json:"is_error"`
}

type SealClusterStateChanged struct {
	State   string `json:"state" validate:"required,min=4,max=20"`
	Region  string `json:"region" validate:"required,max=50"`
	Message string `json:"message" validate:"max=500"`
	IsError bool   `json:"is_error"`
}

type UnsealClusterStateChanged struct {
	State   string `json:"state" validate:"required,min=4,max=20"`
	Region  string `json:"region" validate:"required,max=50"`
	Message string `json:"message" validate:"max=500"`
	IsError bool   `json:"is_error"`
}

type UpdateClusterStateChanged struct {
	ArxId   common.ArxId `json:"cluster_id"`
	State   string       `json:"state" validate:"required,min=4,max=20"`
	Region  string       `json:"region" validate:"required,max=50"`
	Message string       `json:"message" validate:"max=500"`
	IsError bool         `json:"is_error"`
}

type RotateClusterKeysStateChanged struct {
	ArxId   common.ArxId `json:"cluster_id"`
	State   string       `json:"state" validate:"required,min=4,max=20"`
	Region  string       `json:"region" validate:"required,max=50"`
	Message string       `json:"message" validate:"max=500"`
	IsError bool         `json:"is_error"`
}

type APIKeyStateChanged struct {
	KeyID                 string `json:"key_id" validate:"required"`
	AccessKey             string `json:"access_key" validate:"required"`
	ServerSigningKey      string `json:"apikey_signing_key,omitempty"`
	ServerCryptoAccessKey string `json:"apikey_crypto_key,omitempty"`
	ClientKey             string `json:"apikey_client_key,omitempty"`
	ClusterID             string `json:"cluster_id" validate:"required"`
	ApplicationID         string `json:"application_id" validate:"required"`
	State                 string `json:"state" validate:"required,min=4,max=20"`
	Region                string `json:"region" validate:"required,max=50"`
	Message               string `json:"message" validate:"max=500"`
	ArxUrl                string `json:"arx_url"`
	IsError               bool   `json:"is_error"`
}

type DisableApplicationStateChanged struct {
	ApplicationID string `json:"application_id" validate:"required"`
	State         string `json:"state" validate:"required,min=4,max=20"`
	Region        string `json:"region" validate:"required,max=50"`
	Message       string `json:"message" validate:"max=500"`
	IsError       bool   `json:"is_error"`
}

type EnableApplicationStateChanged struct {
	ApplicationID string `json:"application_id" validate:"required"`
	State         string `json:"state" validate:"required,min=4,max=20"`
	Region        string `json:"region" validate:"required,max=50"`
	Message       string `json:"message" validate:"max=500"`
	IsError       bool   `json:"is_error"`
}

type ArchiveApplicationStateChanged struct {
	ApplicationID string `json:"application_id" validate:"required"`
	State         string `json:"state" validate:"required,min=4,max=20"`
	Region        string `json:"region" validate:"required,max=50"`
	Message       string `json:"message" validate:"max=500"`
	IsError       bool   `json:"is_error"`
}
