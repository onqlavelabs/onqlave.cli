package application

import (
	"time"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/acl"
	"github.com/onqlavelabs/onqlave.cli/core/contracts/common"
)

type RequestApplication struct {
	Name        string `json:"name" validate:"required,max=150"`
	Description string `json:"description" validate:"max=500"`
	Technology  string `json:"technology" validate:"required,max=20"`
	Owner       string `json:"owner" validate:"required,max=150"`
	Cors        []Cors `json:"cors" validate:"max=10"`
}

type Application struct {
	ACL         acl.ACL              `json:"acl"`
	ID          common.ApplicationId `json:"application_id" validate:"required"`
	Name        string               `json:"name" validate:"required,max=150"`
	Description string               `json:"description" validate:"required,max=500"`
	Technology  string               `json:"technology" validate:"required,max=20"`
	Owner       string               `json:"owner" validate:"required,max=150"`
	APIKeys     int                  `json:"api_keys"`
	Cors        []Cors               `json:"cors" validate:"max=10"`
	Status      string               `json:"status" validate:"required"`
}

type Applications struct {
	ACL          acl.ACL       `json:"acl"`
	Applications []Application `json:"applications"`
	Models       Technologies  `json:"model"`
	Statistics   Statistics    `json:"statistics"`
}

type Technologies struct {
	Technologies []Technology `json:"technologies" validate:"required"`
}

type Statistics struct {
	Total    int16 `json:"total_applications"`
	Sealed   int16 `json:"sealed_applications"`
	Archived int16 `json:"archived_applications"`
}

type Cors struct {
	Address string `json:"address" validate:"required,max=300,url"`
}

type Status struct {
	ID        common.ApplicationId `json:"application_id" validate:"required"`
	State     string               `json:"data"`
	Message   string               `json:"message"`
	Operation string               `json:"operation"`
	IsError   bool                 `json:"is_error"`
	UpdatedAt time.Time            `json:"update_time"`
}

type Technology struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Cors        bool   `json:"cors" validate:"required"`
	Order       int16  `json:"order" validate:"required"`
	Icon        string `json:"icon" validate:"required"`
	Enable      bool   `json:"enable" validate:"required"`
	IsDefault   bool   `json:"is_default" validate:"required"`
}

type ClientType struct {
	Id    int    `json:"id" validate:"required"`
	Title string `json:"title" validate:"required"`
}
