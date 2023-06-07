package domainevents

import "github.com/google/uuid"

type BillingAccountCreated struct {
	EventId          uuid.UUID `json:"event_id" validate:"required,max=100"`
	TenantID         string    `json:"tenant_id"  validate:"required,min=10,max=200"`
	SubscriptionType string    `json:"subscription_type" validate:"required,max=100"`
}

func (d *BillingAccountCreated) SetEventId(id uuid.UUID) {
	d.EventId = id
}

func (d *BillingAccountCreated) GetEventId() uuid.UUID {
	return d.EventId
}

func (d *BillingAccountCreated) Topic() string {
	return "BillingAccountCreated"
}

func (d *BillingAccountCreated) Metadata() map[string]string {
	return nil
}
