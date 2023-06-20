package errors

import (
	"fmt"
)

type InternalError struct {
	Scope         string `json:"scope,omitempty"`
	Key           string `json:"key,omitempty"`
	CorrelationID string `json:"correlation_id,omitempty"`
	Message       string `json:"message,omitempty"`
	BaseError     error
}

func (e *InternalError) Error() string {
	if e.Scope == ScopeCLI {
		return e.Message
	}
	return fmt.Sprintf("%s error - key:%s", e.Scope, e.Key)
}

func NewInternalError(scope, key, cId, message string, base error) *InternalError {
	return &InternalError{
		Scope:         scope,
		Key:           key,
		CorrelationID: cId,
		Message:       message,
		BaseError:     base,
	}
}

func NewHandlerError(key, cId, message string, err error) *InternalError {
	intErr, ok := err.(*InternalError)
	if ok && intErr != nil {
		if intErr.Key != "" {
			key = intErr.Key
		}

		if intErr.Message != "" {
			message = intErr.Message
		}
	}

	return NewInternalError(ScopeHandler, key, cId, message, err)
}

func NewServiceError(key, cId, message string, err error) *InternalError {
	intErr, ok := err.(*InternalError)
	if ok && intErr != nil {
		if intErr.Key != "" {
			key = intErr.Key
		}

		if intErr.Message != "" {
			message = intErr.Message
		}
	}

	return NewInternalError(ScopeService, key, cId, message, err)
}

func NewPackageError(key string, err error) *InternalError {
	return NewInternalError(ScopeInternal, key, "", "", err)
}

func NewRepoError(key string, err error) *InternalError {
	return NewInternalError(ScopeRepository, key, "", "", err)
}

func NewDbEmptyResultErr() error {
	return NewRepoError(KeyDbEmptyErr, nil)
}

func NewCLIError(key, errMessage string) error {
	return NewInternalError(ScopeCLI, key, "", errMessage, nil)
}
