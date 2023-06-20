package tenant

import "github.com/onqlavelabs/onqlave.cli/core/contracts/common"

type AddTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type UpdateTenantResponse struct {
	common.BaseErrorResponse
	Tenant TenantInfo `json:"data"`
}

type DisableTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type DeleteTenantResponse struct {
	common.BaseErrorResponse
	TenantId common.TenantId `json:"data"`
}

type GetTenantResponse struct {
	common.BaseErrorResponse
	Tenant TenantInfo `json:"data"`
}

type GetTenantDashboardResponse struct {
	common.BaseErrorResponse
	Dashboard TenantDashboard `json:"data"`
}

type GetTenantsResponse struct {
	common.BaseErrorResponse
	Tenants []Tenant `json:"data"`
}
