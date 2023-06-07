package utils

import (
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	"github.com/labstack/echo/v4"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/domainevents"
)

type Subscriber struct {
	subscriptions []common.Subscription
}

func (s *Subscriber) Register(subscriptions []common.Subscription) {
	if len(subscriptions) > 0 {
		s.subscriptions = append(s.subscriptions, subscriptions...)
	}
}

func (s *Subscriber) InitDapr(group *echo.Group) {
	group.GET("/subscribe", s.DaprHandler)
}

func (s *Subscriber) DaprHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.subscriptions)
}

func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

func InitSubscriptions(groupName string, events []domainevents.DomainEvent) []common.Subscription {
	var subscriptions []common.Subscription
	for _, event := range events {
		subscriptions = append(subscriptions, common.Subscription{
			PubsubName: PubSubMain,
			Topic:      event.Topic(),
			Route:      groupName + "/" + event.Topic(),
		})
	}

	return subscriptions
}
