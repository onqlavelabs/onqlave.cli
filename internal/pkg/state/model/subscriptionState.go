package model

import "encoding/json"

type SubscriptionState struct {
	TenantID string            `json:"tenant_id"`
	Value    []byte            `json:"value"`
	Metadata map[string]string `json:"metadata"`
}

func (p SubscriptionState) Key() string {
	return p.TenantID
}

func (p SubscriptionState) Data() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}

func (p SubscriptionState) MetaData() map[string]string {
	return p.Metadata
}
