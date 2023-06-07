package domainevents

import "github.com/google/uuid"

type TenantSubscriptionUpgraded struct {
	EventId          uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID         string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	SubscriptionType string    `json:"subscription_type" validate:"required,max=100"`
}

func (d *TenantSubscriptionUpgraded) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *TenantSubscriptionUpgraded) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *TenantSubscriptionUpgraded) Topic() string {
	return "TenantSubscriptionUpgraded"
}

func (d *TenantSubscriptionUpgraded) Metadata() map[string]string {
	return nil
}
