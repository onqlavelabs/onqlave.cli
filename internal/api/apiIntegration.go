package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type RegisterationRequest struct {
	Request Registration `json:"registration" validate:"required"`
}

type Registration struct {
	Email      string `json:"email_address" validate:"email,required"`
	Operation  string `json:"operation" validate:"required"`
	TenantName string `json:"tenant_name" validate:"required,min=4,max=100"`
}

type RegistrationStatusRequest struct {
	Request RegistrationToken `json:"request" validate:"required"`
}

type RegistrationToken struct {
	Token string `json:"token" validate:"required"`
}

type APIIntegrationService struct {
	opts APIIntegrationServiceOptions
}

type APIIntegrationServiceOptions struct {
	Ctx context.Context
}

type BaseErrorResponse struct {
	Error struct {
		ErrorCode    int32  `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"error"`
}
type RegistrationResponse struct {
	BaseErrorResponse
	Id RegistrationID `json:"data"`
}

type RegistrationID struct {
	Id string `json:"id" validate:"required"`
}

type RegistrationStatusResponse struct {
	BaseErrorResponse
	Status RegistrationStatus `json:"data"`
}

type RegistrationStatus struct {
	Code       string `json:"status"`
	Message    string `json:"message"`
	TenantName string `json:"tenant_name"`
	Token      string `json:"token"`
	TenantID   string `json:"tenant_id"`
}

type State int64
type Operations int64

const (
	Login Operations = iota
	Signup
)

const (
	Initiated State = iota
	Waiting
	Pending
	Completed
	Expired
)

func (o Operations) String() string {
	switch o {
	case Login:
		return "login"
	case Signup:
		return "signup"
	}
	return "unknown"
}

func (s State) String() string {
	switch s {
	case Initiated: // initially requested from client
		return "initiated"
	case Waiting:
		return "waiting"
	case Pending: // email is sent to email address, waiting for signup
		return "pending"
	case Completed: //signup is completed
		return "completed"
	case Expired: //request is not actioned in timely manner so it is expired
		return "expired"
	}
	return "unknown"
}

type Tenant struct {
	Name        string `json:"name" validate:"required,min=4,max=100"`
	Description string `json:"description" validate:"max=500"`
	Id          string `json:"tenant_id" validate:"required"`
	Disable     bool   `json:"disable"`
	OwnerEmail  string `json:"owner_email" validate:"email,required"`
	RequestId   string `json:"request_id" validate:"required"`
}

func NewAPIIntegrationService(opts APIIntegrationServiceOptions) *APIIntegrationService {
	return &APIIntegrationService{
		opts: opts,
	}
}

func _post[Response any, Request any](apiBase string, request Request) (*Response, error) {
	client := &http.Client{}
	payload, _ := json.Marshal(&request)

	req, err := http.NewRequest(http.MethodPost, apiBase, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("ONQLAVE-API-KEY", "")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(bodyBytes))
	}
	var responseObject Response
	json.Unmarshal(bodyBytes, &responseObject)
	return &responseObject, nil
}

func _get[T any](apiBase string) (*T, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiBase, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("ONQLAVE-API-KEY", "")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", viper.GetString("auth_key")))
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(bodyBytes))
	}
	var responseObject T
	json.Unmarshal(bodyBytes, &responseObject)
	return &responseObject, nil
}

func (s *APIIntegrationService) SendSignupInvitation(emailAddress string, tenantName string) (string, error) {
	baseUrl := viper.Get("api_base_url")
	registrationUrl := fmt.Sprintf("%s:%d/registrations/new", baseUrl, 8081)
	request := RegisterationRequest{
		Request: Registration{
			Email: emailAddress, Operation: Signup.String(), TenantName: tenantName,
		},
	}
	response, err := _post[RegistrationResponse](registrationUrl, request)
	if err != nil {
		return "", err
	}
	return response.Id.Id, nil
}

type GetTenantResponse struct {
	BaseErrorResponse
	Tenant Tenant `json:"data"`
}

func (s *APIIntegrationService) GetTenant() (map[string]interface{}, error) {
	baseUrl := viper.Get("api_base_url")
	tenantId := viper.Get("tenant_id")
	tenantUrl := fmt.Sprintf("%s:%d/tenants/%s", baseUrl, 8083, tenantId)

	response, err := _get[map[string]interface{}](tenantUrl)
	if err != nil {
		return nil, err
	}
	return *response, nil
}

func (s *APIIntegrationService) SendLoginInvitation(emailAddress string, tenantName string) (string, error) {
	baseUrl := viper.Get("api_base_url")
	registrationUrl := fmt.Sprintf("%s:%d/registrations/new", baseUrl, 8081)
	request := RegisterationRequest{
		Request: Registration{
			Email: emailAddress, Operation: Login.String(), TenantName: tenantName,
		},
	}
	response, err := _post[RegistrationResponse](registrationUrl, request)
	if err != nil {
		return "", err
	}
	return response.Id.Id, nil
}

func (s *APIIntegrationService) GetSignupOperationStatus(token string) (*APIIntegrationServiceOperationResult, error) {
	baseUrl := viper.Get("api_base_url")
	registrationUrl := fmt.Sprintf("%s:%d/registrations/status", baseUrl, 8081)
	request := RegistrationStatusRequest{
		Request: RegistrationToken{
			Token: token,
		},
	}
	response, err := _post[RegistrationStatusResponse](registrationUrl, request)
	if err != nil {
		return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for signup completion."}, err
	}
	if response.Status.Code == Completed.String() {
		return &APIIntegrationServiceOperationResult{Done: true, Result: "Waiting for signup completion."}, nil
	} else if response.Status.Code == Expired.String() {
		return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for signup completion."}, fmt.Errorf("request timed out")
	}
	return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for signup completion."}, nil
}

func (s *APIIntegrationService) GetLoginOperationStatus(token string) (*APIIntegrationServiceOperationResult, string, string, error) {
	baseUrl := viper.Get("api_base_url")
	registrationUrl := fmt.Sprintf("%s:%d/registrations/status", baseUrl, 8081)
	request := RegistrationStatusRequest{
		Request: RegistrationToken{
			Token: token,
		},
	}
	var authToken string
	var tenantID string
	response, err := _post[RegistrationStatusResponse](registrationUrl, request)
	if err != nil {
		return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for login completion."}, authToken, tenantID, err
	}
	if response.Status.Code == Completed.String() {
		return &APIIntegrationServiceOperationResult{Done: true, Result: "Waiting for login completion."}, response.Status.Token, response.Status.TenantID, nil
	} else if response.Status.Code == Expired.String() {
		return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for login completion."}, authToken, tenantID, fmt.Errorf("request timed out")
	}
	return &APIIntegrationServiceOperationResult{Done: false, Result: "Waiting for login completion."}, authToken, tenantID, nil
}

type APIIntegrationServiceOperationResult struct {
	Done   bool
	Result any
}
