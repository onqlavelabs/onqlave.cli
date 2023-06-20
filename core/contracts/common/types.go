package common

import (
	"regexp"
)

type ArxId string
type ApplicationId string
type TenantId string
type UserId string
type UserEmail string

func (id *ArxId) Valid() bool {
	return true
}

func (id *ApplicationId) Valid() bool {
	return true
}

func (id *TenantId) Valid() bool {
	pattern := regexp.MustCompile(`--`)
	result := pattern.Split(string(*id), -1)
	return len(result) == 2
}
