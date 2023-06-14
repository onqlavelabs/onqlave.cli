package models

import (
	"github.com/onqlavelabs/onqlave.cli/core/contracts"
)

type UserList struct {
	Users []contracts.User `json:"users"`
}
