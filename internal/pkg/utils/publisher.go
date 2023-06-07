package utils

import (
	"context"
	"encoding/json"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/google/uuid"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/domainevents"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/errors"
)

const (
	PubSubMain = "pubsub"
)

type daprPublisher[T domainevents.DomainEvent] struct {
}

type DAPRPublisherService[T domainevents.DomainEvent] interface {
	Publish(context context.Context, event T) error
}

func (s *daprPublisher[T]) Publish(context context.Context, event T) error {
	client, err := dapr.NewClient()
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	if client == nil {
		err := fmt.Errorf("dapr is not running properly. please check the configurations")
		return errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	event.SetEventId(uuid.New())
	data, err := json.Marshal(event)
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	if err := client.PublishEvent(context, PubSubMain, event.Topic(), data); err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprPublishErr, err)
	}

	return nil
}

func NewDAPRPublisherService[T domainevents.DomainEvent]() DAPRPublisherService[T] {
	return &daprPublisher[T]{}
}
