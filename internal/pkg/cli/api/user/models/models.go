package models

import "github.com/onqlavelabs/onqlave.cli/internal/pkg/tenant/contracts"

type UserList struct {
	Users []contracts.User `json:"users"`
}
