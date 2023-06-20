package common

import (
	"regexp"
)

type BaseErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code          int           `json:"code" yaml:"code"`                     // Code represent codes in response
	Status        string        `json:"status" yaml:"status"`                 // Status represent string value of code
	Message       string        `json:"message" yaml:"message"`               // Message represent detail message
	CorrelationID string        `json:"correlation_id" yaml:"correlation_id"` // The RequestId that's also set in the header
	Details       []interface{} `json:"details" yaml:"details"`               // Details is a list of details in any types in string
}

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
