package api

import (
	"context"
	"fmt"
	"time"
)

type TenantService struct {
	opts    TenantServiceOptions
	counter int
}

type TenantServiceOptions struct {
	Ctx       context.Context
	Iteration int
}

func NewTenantService(opts TenantServiceOptions) *TenantService {
	return &TenantService{
		opts: opts,
	}
}

func (t *TenantService) GetSignupLink() (string, error) {
	link := "http://localhost:8080/8jhgjsdhfgjhgsdf878789sfjhjhksfkhjkhsfjhjkhsfjkhsdfhskfhjdfkjh"
	return link, nil
}

func (t *TenantService) SendSignupInvitation(emailAddress string, tenantName string) error {
	return nil
}

func (t *TenantService) GetSignupOperationStatus() (result *TenantOperationResult, err error) {
	time.Sleep(5 * time.Second)
	t.counter++
	if t.counter < t.opts.Iteration {
		return &TenantOperationResult{Done: false, Result: fmt.Sprintf("Waiting for signup completion. Operation in progress..%d", t.counter)}, nil
	}
	return &TenantOperationResult{Done: true, Result: nil}, nil
}

type TenantOperationResult struct {
	Done   bool
	Result any
}
