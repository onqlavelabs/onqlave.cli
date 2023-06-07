package common

type OnqlaveIdHeader string
type OnqlaveApiKeyHeader string
type OnqlaveVersionHeader string
type OnqlaveRouteHeader string
type OnqlaveSignatureHeader string
type OnqlaveUserHeader string

const (
	OnqlaveId        OnqlaveIdHeader        = "ONQLAVE-ID"
	OnqlaveApiKey    OnqlaveApiKeyHeader    = "ONQLAVE-API-KEY"
	OnqlaveContext   string                 = "ONQLAVE-CONTEXT"
	OnqlaveVersion   OnqlaveVersionHeader   = "ONQLAVE-VERSION"
	OnqlaveRoute     OnqlaveRouteHeader     = "ONQLAVE-ROUTE"
	OnqlaveSignature OnqlaveSignatureHeader = "ONQLAVE_SIGANTURE"
	OnqlaveUser      OnqlaveUserHeader      = "ONQLAVE_USER"
)

const (
	ServerType string = "Onqlave/0.1"
)
