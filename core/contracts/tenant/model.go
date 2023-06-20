package tenant

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type TenantId struct {
	Id string `json:"id" validate:"required"`
}

type Tenant struct {
	Id          common.TenantId `json:"tenant_id" validate:"required"`
	Name        string          `json:"name" validate:"required,min=4,max=100"`
	Description string          `json:"description" validate:"max=500"`
	Disable     bool            `json:"disable"`
	OwnerEmail  string          `json:"owner_email" validate:"email,required"`
	RequestId   string          `json:"request_id" validate:"required"`
}

type TenantInfo struct {
	Id         common.TenantId `json:"tenant_id,omitempty"`
	Name       string          `json:"tenant_name" validate:"required,min=4,max=100"`
	Label      string          `json:"tenant_label"  validate:"required"`
	OwnerEmail string          `json:"owner_email,omitempty"`
	CreatedOn  time.Time       `json:"created_on,omitempty"`
	ACL        acl.ACL         `json:"acl,omitempty"`
}

type TenantDashboard struct {
	Insight DasboardInsight  `json:"insight"`
	Events  []DashboardEvent `json:"events"`
}

type DasboardInsight struct {
	NumberOfArx          int `json:"number_of_arx"`
	NumberOfApplications int `json:"number_of_applications"`
	NumberOfApiKeys      int `json:"number_of_apikeys"`
	NumberOfUsers        int `json:"number_of_users"`
}

type DashboardEvent struct {
	Operation string    `json:"operations"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	IsError   bool      `json:"is_error"`
}
