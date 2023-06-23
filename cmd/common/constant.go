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
	prefix               = `onqlave`
	help                 = `Usage:
	{{if .Runnable}}
		{{.UseLine}}
	{{end}}
	{{if .HasAvailableSubCommands}}
		{{.CommandPath}} [command]
	{{end}}
	{{if gt (len .Aliases) 0}}
  	Aliases:
		{{.NameAndAliases}}
	{{end}}
  {{if .HasExample}}
  Examples:
	{{.Example}}
  {{end}}
  {{if .HasAvailableSubCommands}}
  Available Commands:
	{{range .Commands}}
		{{if (or .IsAvailableCommand (eq .Name "help"))}}
			{{rpad .Name .NamePadding }} {{.Short}}
		{{end}}
	{{end}}
  {{end}}
  {{if .HasAvailableLocalFlags}}
  Flags:
  {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}
  {{end}}
  {{if .HasAvailableInheritedFlags}}
  Global Flags:
  {{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}
  {{end}}
  {{if .HasHelpSubCommands}}
  Additional help topics:
	{{range .Commands}}
		{{if .IsAdditionalHelpTopicCommand}}
		{{rpad .CommandPath .CommandPathPadding}} {{.Short}}
		{{end}}
	{{end}}
  {{end}}
  {{if .HasAvailableSubCommands}}
  Use "{{.CommandPath}} [command] --help" for more information about a command.
  {{end}}
  `
)
