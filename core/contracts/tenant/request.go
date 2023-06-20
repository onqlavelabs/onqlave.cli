package tenant

type AddTenantRequest struct {
	Tenant Tenant `json:"tenant" validate:"required"`
}

type UpdateTenantRequest struct {
	Tenant TenantInfo `json:"tenant" validate:"required"`
}

type DisableTenantRequest struct {
	Disable *bool `json:"disable" validate:"required"`
}
