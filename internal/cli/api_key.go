package cli

import (
	"fmt"

	"github.com/onqlavelabs/onqlave.cli/core/contracts/api_key"
)

const (
	apiKeyPath = "keys"
)

type ApiKeyService interface {
	List()
	Get(string) (*api_key.APIKey, error)
	Base()
	Create()
	Delete() error
}

type ApiKeyServiceOp struct {
	client *Client
}

func (s *ApiKeyServiceOp) List() {

}

func (s *ApiKeyServiceOp) Get(id string) (*api_key.APIKey, error) {
	path := fmt.Sprintf("%s/%s", apiKeyPath, id)

	resource := new(api_key.DetailResponse)
	err := s.client.Get(path, &resource, nil)
	if err != nil {
		return nil, err
	}

	return &resource.Data, nil
}

func (s *ApiKeyServiceOp) Base() {

}

func (s *ApiKeyServiceOp) Create() {

}

func (s *ApiKeyServiceOp) Delete() error {
	return nil
}
