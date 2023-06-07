package utils

import (
	daprCommon "github.com/dapr/go-sdk/service/common"
	"github.com/labstack/echo/v4"
)

type HandlerSetup interface {
	SetupRoutes(groupName string, group *echo.Group) []daprCommon.Subscription
}

type BindingHandlerSetup interface {
	SetupBinding() error
}
