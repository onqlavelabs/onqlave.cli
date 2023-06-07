package model

import (
	"encoding/json"
)

type SubscriptionTypeProductState struct {
	Value    []byte            `json:"value"`
	Metadata map[string]string `json:"metadata"`
}

func (p SubscriptionTypeProductState) Key() string {
	return "subscription_type_product_state"
}

func (p SubscriptionTypeProductState) Data() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}

func (p SubscriptionTypeProductState) MetaData() map[string]string {
	return p.Metadata
}
