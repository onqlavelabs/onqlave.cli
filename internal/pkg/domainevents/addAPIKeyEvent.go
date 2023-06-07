package domainevents

import (
	"time"

	"github.com/google/uuid"

	"github.com/onqlavelabs/onqlave.cli/internal/app/tenant/apikey/enums"
)

type APIKeyStatusChange struct {
	EventId               uuid.UUID          `json:"event_id" validate:"required,max=100"`
	TenantID              string             `json:"tenant_id" validate:"required"`
	KeyID                 string             `json:"id" validate:"required"`
	ApplicationID         string             `json:"application_id" validate:"required"`
	ClusterID             string             `json:"cluster_id" validate:"required"`
	AccessKey             string             `json:"access_key" validate:"min=10"`
	ClientKey             string             `json:"client_key" validate:"min=10"`
	ServerSigningKey      string             `json:"server_signing_key" validate:"min=10"`
	ServerCryptoAccessKey string             `json:"server_crypto_access_key" validate:"min=10"`
	Status                enums.ApiKeyStatus `json:"status" validate:"required"`
	ProvidedAt            *time.Time         `json:"provided_at" validate:"required"`
	Attributes            map[string]string  `json:"attributes" validate:"max=20"`
	Message               string             `json:"message" validate:"max=500"`
	ArxUrl                string             `json:"arx_url"`
	IsError               bool               `json:"is_error"`
}

func (api *APIKeyStatusChange) SetEventId(id uuid.UUID) {
	api.EventId = id
}

func (api *APIKeyStatusChange) GetEventId() uuid.UUID {
	return api.EventId
}

func (api *APIKeyStatusChange) Topic() string {
	return "APIKeyStatusChanged"
}

func (api *APIKeyStatusChange) Metadata() map[string]string {
	return nil
}
