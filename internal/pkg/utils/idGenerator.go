package utils

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/errors"
)

type IDGenerator interface {
	NewApplicationID() (string, error)
	NewDomainID() (string, error)
	NewTenantID() (string, error)
	NewRegistrationRequestID() (string, error)
	NewEventID() (string, error)
}

type idGenerator struct {
}

func NewIdGenerator() IDGenerator {
	return idGenerator{}
}

func (i idGenerator) NewRegistrationRequestID() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", errors.NewErrGenerateID(err)
	}
	id := fmt.Sprintf("registration--%s", nanoId)
	return id, nil
}

func (i idGenerator) NewEventID() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", errors.NewErrGenerateID(err)
	}
	id := fmt.Sprintf("tenant--%s", nanoId)
	return id, nil
}

func (i idGenerator) NewTenantID() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", errors.NewErrGenerateID(err)
	}
	id := fmt.Sprintf("tenant--%s", nanoId)
	return id, nil
}

func (i idGenerator) NewApplicationID() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", errors.NewErrGenerateID(err)
	}
	id := fmt.Sprintf("app--%s", nanoId)
	return id, nil
}

func (i idGenerator) NewDomainID() (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", errors.NewErrGenerateID(err)
	}
	id := fmt.Sprintf("domain--%s", nanoId)
	return id, nil
}

func NewAPIKeyID() string {
	nanoId, err := gonanoid.New()
	if err != nil {
		return ""
	}
	id := fmt.Sprintf("apikey--%s", nanoId)
	return id
}

func GenerateID(prefix string) (string, error) {
	nanoId, err := gonanoid.New()
	if err != nil {
		return "", nil
	}
	id := fmt.Sprintf("%s--%s", prefix, nanoId)
	return id, nil
}
