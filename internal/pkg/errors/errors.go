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

func NewPackageError(key string, err error) *InternalError {
	return NewInternalError(ScopeInternal, key, "", "", err)
}
