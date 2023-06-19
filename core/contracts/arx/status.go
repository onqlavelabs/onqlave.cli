package contracts

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type ClusterStatus struct {
	ID        common.ArxId   `json:"cluster_id" validate:"required"`
	State     string         `json:"state,omitempty"`
	Message   string         `json:"message,omitempty"`
	Operation string         `json:"operation,omitempty"`
	IsError   bool           `json:"is_error,omitempty"`
	UpdatedAt time.Time      `json:"update_time,omitempty"`
	Cluster   *ClusterDetail `json:"cluster,omitempty"`
	ACL       acl.ACL        `json:"acl"`
}

type ClusterDetail struct {
	TenantID           common.TenantId `json:"tenant_id"`
	Name               string          `json:"name"`
	SpendLimit         uint64          `json:"spend_limit"`
	Description        string          `json:"description"`
	Purpose            string          `json:"purpose"`
	PlanID             string          `json:"plan_id"`
	ProviderID         string          `json:"provider_id"`
	Regions            []string        `json:"regions"`
	EncryptionMethodID string          `json:"encryption_method_id"`
	RotationCycleID    string          `json:"rotation_cycle_id"`
	Owner              string          `json:"owner"`
	IsDefault          bool            `json:"is_default"`
}
