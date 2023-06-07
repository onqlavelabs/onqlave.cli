package utils

import (
	"context"
	"encoding/json"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/errors"
	"github.com/onqlavelabs/onqlave.cli/internal/pkg/state"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	StateStoreName = "statestore"
)

type daprStateStore[T state.State] struct {
}

type DAPRStateStoreService[T state.State] interface {
	Save(ctx context.Context, state T) error
	Get(ctx context.Context, state T) (*T, error)
	Delete(ctx context.Context, state T) error
}

func (s *daprStateStore[T]) Save(ctx context.Context, state T) error {
	client, err := dapr.NewClient()
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	err = client.SaveState(ctx, StateStoreName, state.Key(), state.Data(), state.MetaData())
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprSaveStateErr, err)
	}

	return nil
}

func (s *daprStateStore[T]) Get(ctx context.Context, state T) (*T, error) {
	client, err := dapr.NewClient()
	if err != nil {
		return nil, errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	st, err := client.GetState(ctx, StateStoreName, state.Key(), map[string]string{})
	if err != nil {
		return nil, errors.NewPackageError(errors.KeyInternalDaprGetStateErr, err)
	}
	if st.Value == nil {
		return nil, errors.NewPackageError(errors.KeyInternalDaprStateNotFoundErr, nil)
	}

	var result T
	if err := json.Unmarshal(st.Value, &result); err != nil {
		return nil, errors.NewPackageError(errors.KeyInternalDaprStateNotFoundErr, nil)
	}

	return &result, nil
}

func (s *daprStateStore[T]) Delete(ctx context.Context, state T) error {
	client, err := dapr.NewClient()
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprInitErr, err)
	}

	err = client.DeleteState(ctx, StateStoreName, state.Key(), map[string]string{})
	if err != nil {
		return errors.NewPackageError(errors.KeyInternalDaprDeleteStateErr, err)
	}

	return nil
}

func NewDAPRStateStoreService[T state.State]() DAPRStateStoreService[T] {
	return &daprStateStore[T]{}
}
