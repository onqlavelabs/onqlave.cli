package common

import (
	"github.com/labstack/echo/v4"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/auth/key"
)

type onqlaveOperationContext struct {
	onqlaveId OnqlaveIdHeader
	apiKey    OnqlaveApiKeyHeader
	routeId   OnqlaveRouteHeader
	signature OnqlaveSignatureHeader
	version   OnqlaveVersionHeader
	user      *key.Token
}

type IOnqlaveOperationContext interface {
	SetApiKey(apiKey OnqlaveApiKeyHeader)
	SetOnqlaveId(onqlaveId OnqlaveIdHeader)
	SetRouteId(tenantId OnqlaveRouteHeader)
	SetVersion(version OnqlaveVersionHeader)
	SetSignature(signature OnqlaveSignatureHeader)
	GetApiKey() OnqlaveApiKeyHeader
	GetOnqlaveId() OnqlaveIdHeader
	GetRouteId() OnqlaveRouteHeader
	GetSignature() OnqlaveSignatureHeader
	GetVersion() OnqlaveVersionHeader
	GetUser() *key.Token
	SetUser(token *key.Token)
}

func (c *onqlaveOperationContext) SetApiKey(apiKey OnqlaveApiKeyHeader) {
	c.apiKey = apiKey
}

func (c *onqlaveOperationContext) SetOnqlaveId(onqlaveId OnqlaveIdHeader) {
	c.onqlaveId = onqlaveId
}

func (c *onqlaveOperationContext) GetApiKey() OnqlaveApiKeyHeader {
	return c.apiKey
}

func (c *onqlaveOperationContext) GetRouteId() OnqlaveRouteHeader {
	return c.routeId
}

func (c *onqlaveOperationContext) SetRouteId(routeId OnqlaveRouteHeader) {
	c.routeId = routeId
}

func (c *onqlaveOperationContext) GetOnqlaveId() OnqlaveIdHeader {
	return c.onqlaveId
}

func (c *onqlaveOperationContext) SetSignature(signature OnqlaveSignatureHeader) {
	c.signature = signature
}

func (c *onqlaveOperationContext) SetVersion(version OnqlaveVersionHeader) {
	c.version = version
}

func (c *onqlaveOperationContext) GetSignature() OnqlaveSignatureHeader {
	return c.signature
}

func (c *onqlaveOperationContext) GetVersion() OnqlaveVersionHeader {
	return c.version
}

func (c *onqlaveOperationContext) GetUser() *key.Token {
	return c.user
}

func (c *onqlaveOperationContext) SetUser(token *key.Token) {
	c.user = token
}

func GetOnqlaveContext(c echo.Context) (bool, IOnqlaveOperationContext) {
	newObject := false
	ctx := c.Get(OnqlaveContext)
	if ctx == nil {
		ctx = &onqlaveOperationContext{}
		newObject = true
	}
	context := ctx.(IOnqlaveOperationContext)
	return newObject, context
}
