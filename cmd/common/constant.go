package common

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"

	BaseUrlDev  = "https://dev.onqlave.io"
	BaseUrlProd = "https://api.onqlave.com"

	ConfigFile     = "config"
	configDir      = "/.config/onqlave/"
	ConfigTypeJson = "json"

	FlagApiBaseUrl = "api_base_url"
	FlagAuthKey    = "auth_key"
	FlagConfigPath = "config_path"
	FlagEnv        = "env"
	FlagJson       = "json"
	FlagTenantID   = "tenant_id"

	Valid   = 30
	Version = `0.0.1`

	ResourceArx         = "Arx"
	ResourceApplication = "Application"
	ResourceKey         = "API Key"
	ResourceTenant      = "Tenant"
	ResourceUser        = "User"

	ActionCreated    = "created"
	ActionUpdated    = "updated"
	ActionDeleted    = "deleted"
	ActionArchived   = "archived"
	ActionEnabled    = "enabled"
	ActionDisabled   = "disabled"
	ActionSealed     = "sealed"
	ActionUnsealed   = "unsealed"
	ActionSetDefault = "set default"

	TableViewWidth       = 100
	TableViewHeight      = 10
	TableViewMaxColWidth = 35
	TableViewMinColWidth = 11

	OSWindows = "windows"
	OSLinux   = "linux"
	OSDarwin  = "darwin"
)
