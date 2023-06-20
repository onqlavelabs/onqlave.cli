package contracts

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type NewArx struct {
	Name             string   `json:"name" validate:"required"`
	SpendLimit       *uint64  `json:"spend_limit" validate:"required"`
	Purpose          string   `json:"purpose" validate:"required,max=50"`
	Plan             string   `json:"plan" validate:"required,max=50"`
	Provider         string   `json:"provider" validate:"required,max=50"`
	Regions          []string `json:"regions" validate:"required,max=5"`
	EncryptionMethod string   `json:"encryption_method" validate:"required,max=50"`
	RotationCycle    string   `json:"rotation_cycle" validate:"required,max=50"`
	Owner            string   `json:"owner" validate:"required,max=150"`
	IsDefault        bool     `json:"is_default"`
}

type UpdateArx struct {
	ID            common.ArxId `json:"id" validate:"required"`
	Name          string       `json:"name" validate:"required"`
	SpendLimit    *uint64      `json:"spend_limit" validate:"required"`
	Regions       []string     `json:"regions" validate:"required,max=5"`
	RotationCycle string       `json:"rotation_cycle" validate:"required,max=50"`
	Owner         string       `json:"owner" validate:"required,max=150"`
	IsDefault     *bool        `json:"is_default" validate:"required"`
}

type Status struct {
	ID        common.ArxId `json:"cluster_id" validate:"required"`
	State     string       `json:"state,omitempty"`
	Message   string       `json:"message,omitempty"`
	Operation string       `json:"operation,omitempty"`
	IsError   bool         `json:"is_error,omitempty"`
	UpdatedAt time.Time    `json:"update_time,omitempty"`
	Cluster   *Detail      `json:"cluster,omitempty"`
	ACL       acl.ACL      `json:"acl"`
}

type Detail struct {
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

type ExistingWithDetail struct {
	ACL                 acl.ACL      `json:"acl"`
	ID                  common.ArxId `json:"id" validate:"required"`
	Name                string       `json:"name" validate:"required"`
	SpendLimit          uint64       `json:"spend_limit"`
	Description         string       `json:"description" validate:"required,min=4,max=500"`
	Purpose             string       `json:"purpose" validate:"required,max=50"`
	Plan                string       `json:"plan" validate:"required,max=50"`
	Provider            string       `json:"provider" validate:"required,max=50"`
	Regions             []string     `json:"regions" validate:"required,max=5"`
	AvailabilityMessage string       `json:"availability_message"`
	EncryptionMethod    string       `json:"encryption_method" validate:"required,max=50"`
	RotationCycle       string       `json:"rotation_cycle" validate:"required,max=50"`
	Owner               string       `json:"owner" validate:"required,max=150"`
	IsDefault           bool         `json:"is_default"`
	Status              string       `json:"status" validate:"required,max=50"`
}

type EncryptionRotationCycle struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty" validate:"required"`
	IsDefault   *bool  `json:"is_default,omitempty" validate:"required"`
	Enable      *bool  `json:"enable,omitempty" validate:"required"`
	Order       *uint8 `json:"order,omitempty" validate:"required"`
}

type GetListResponse struct {
	ACL      acl.ACL              `json:"acl"`
	Clusters []ExistingWithDetail `json:"clusters"`
	Models   BaseInfo             `json:"model"`
	Insights Insights             `json:"insights"`
}

type BaseInfo struct {
	Purposes          []Purpose                 `json:"purposes" validate:"required"`
	Plans             []Plan                    `json:"plans" validate:"required"`
	EncryptionMethods []EncryptionMethod        `json:"encryption_methods" validate:"required"`
	RotationCycles    []EncryptionRotationCycle `json:"rotation_cycles" validate:"required"`
	Providers         []Provider                `json:"providers" validate:"required"`
}

type Plan struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDefault   *bool  `json:"is_default" validate:"required"`
	Enable      *bool  `json:"enable,omitempty" validate:"required"`
	Order       *uint8 `json:"order,omitempty" validate:"required"`
	Icon        string `json:"icon,omitempty" validate:"required"`
}

type Provider struct {
	ID          string   `json:"id" validate:"required"`
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	IsDefault   *bool    `json:"is_default,omitempty" validate:"required"`
	Enable      *bool    `json:"enable,omitempty" validate:"required"`
	Order       *uint8   `json:"order,omitempty" validate:"required"`
	Image       string   `json:"image" validate:"required"`
	Regions     []Region `json:"regions,omitempty" validate:"required"`
}

type Purpose struct {
	ID        string `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	IsDefault *bool  `json:"is_default,omitempty" validate:"required"`
	Enable    *bool  `json:"enable,omitempty" validate:"required"`
	Order     *uint8 `json:"order,omitempty" validate:"required"`
}

type RegionOptimisation struct {
	Value   int32  `json:"value" validate:"required"`
	Message string `json:"message" validate:"required,max=50"`
}

type Region struct {
	ID           string             `json:"id" validate:"required"`
	Name         string             `json:"name" validate:"required"`
	IsDefault    *bool              `json:"is_default,omitempty" validate:"required"`
	Enable       *bool              `json:"enable,omitempty" validate:"required"`
	Order        *uint8             `json:"order,omitempty" validate:"required"`
	Icon         string             `json:"icon" validate:"required"`
	Optimisation RegionOptimisation `json:"optimisation" validate:"required"`
}

type EncryptionMethod struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	IsDefault   *bool  `json:"is_default,omitempty" validate:"required"`
	Enable      *bool  `json:"enable,omitempty" validate:"required"`
	Order       *uint8 `json:"order,omitempty" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
}

type Insights struct {
	TotalCluster int `json:"total_clusters"`
	TotalActive  int `json:"total_active"`
	TotalSealed  int `json:"total_sealed"`
}
