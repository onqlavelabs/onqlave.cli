package application

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type Application struct {
	Name        string            `json:"name" validate:"required,max=150"`
	Description string            `json:"description" validate:"max=500"`
	Technology  string            `json:"technology" validate:"required,max=20"`
	Owner       string            `json:"owner" validate:"required,max=150"`
	Cors        []ApplicationCors `json:"cors" validate:"max=10"`
}

type ApplicationCors struct {
	Address string `json:"address" validate:"required,max=300,url"`
}

type ExistingApplicationWithDetails struct {
	ACL         acl.ACL              `json:"acl"`
	ID          common.ApplicationId `json:"application_id" validate:"required"`
	Name        string               `json:"name" validate:"required,max=150"`
	Description string               `json:"description" validate:"required,max=500"`
	Technology  string               `json:"technology" validate:"required,max=20"`
	Owner       string               `json:"owner" validate:"required,max=150"`
	APIKeys     int                  `json:"api_keys"`
	Cors        []ApplicationCors    `json:"cors" validate:"max=10"`
	Status      string               `json:"status" validate:"required"`
}

type ApplicationModelWrapper struct {
	Technologies []ApplicationTechnology `json:"technologies" validate:"required"`
}

type ApplicationStatistics struct {
	Total    int16 `json:"total_applications"`
	Sealed   int16 `json:"sealed_applications"`
	Archived int16 `json:"archived_applications"`
}

type ApplicationStatus struct {
	ID        common.ApplicationId `json:"application_id" validate:"required"`
	State     string               `json:"data"`
	Message   string               `json:"message"`
	Operation string               `json:"operation"`
	IsError   bool                 `json:"is_error"`
	UpdatedAt time.Time            `json:"update_time"`
}

type ApplicationTechnology struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Cors        bool   `json:"cors" validate:"required"`
	Order       int16  `json:"order" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Enable      bool   `json:"enable" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
}

type GetApplications struct {
	ACL          acl.ACL                          `json:"acl"`
	Applications []ExistingApplicationWithDetails `json:"applications"`
	Models       ApplicationModelWrapper          `json:"model"`
	Statistics   ApplicationStatistics            `json:"statistics"`
}
