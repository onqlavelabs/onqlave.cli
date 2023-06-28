package common

import (
	"context"

	"github.com/onqlavelabs/onqlave.cli/internal/utils"
	"github.com/onqlavelabs/onqlave.core/errors"
)

type Token struct {
	AuthTime int64                  `json:"auth_time"`
	Issuer   string                 `json:"iss"`
	Audience string                 `json:"aud"`
	Expires  int64                  `json:"exp"`
	IssuedAt int64                  `json:"iat"`
	Subject  string                 `json:"sub,omitempty"`
	UID      string                 `json:"uid,omitempty"`
	Claims   map[string]interface{} `json:"claims,omitempty"`
}

func (t *Token) GetRoles(ctx context.Context) ([]string, error) {
	if _, ok := t.Claims["roles"]; !ok {
		return nil, errors.NewCLIError(errors.KeyCLIInvalidValue, utils.BoldStyle.Render("Invalid roles"))
	}
	result := t.Claims["roles"].([]interface{})
	roles := make([]string, len(result))
	for idx, role := range result {
		roles[idx] = role.(string)
	}
	return roles, nil
}

func (t *Token) GetUserId() string {
	return t.UID
}

func (t *Token) GetTenantId() string {
	if _, ok := t.Claims["tenant_id"]; !ok {
		return ""
	}
	return t.Claims["tenant_id"].(string)
}
