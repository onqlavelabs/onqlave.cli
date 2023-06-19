package common

import (
	"context"
	"fmt"

	"github.com/onqlavelabs/onqlave.cli/internal/pkg/errors"
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
		return nil, errors.NewPackageError("", fmt.Errorf(""))
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
