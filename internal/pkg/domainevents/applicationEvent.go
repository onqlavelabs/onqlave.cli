package domainevents

import (
	"time"

	"github.com/google/uuid"

	"github.com/onqlavelabs/onqlave.cli/internal/app/tenant/apikey/enums"
)

type ApplicationStatusChange struct {
	EventId       uuid.UUID          `json:"event_id" validate:"required,max=100"`
	ApplicationID string             `json:"application_id" validate:"required"`
	ClusterID     string             `json:"cluster_id" validate:"required"`
	Status        enums.ApiKeyStatus `json:"status" validate:"required"`
	ProvidedAt    *time.Time         `json:"provided_at" validate:"required"`
	Attributes    map[string]string  `json:"attributes" validate:"max=20"`
	Message       string             `json:"message" validate:"max=500"`
	IsError       bool               `json:"is_error"`
}

func (d *ApplicationStatusChange) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *ApplicationStatusChange) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *ApplicationStatusChange) Topic() string {
	return "ApplicationStateChanged"
}

func (d *ApplicationStatusChange) Metadata() map[string]string {
	return nil
}
