package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/viper"
)

type Client struct {
	// HTTP client used to communicate with the API.
	Client *http.Client

	// Base URL for API requests.
	baseURL *url.URL

	tenantId    string
	apiKey      string
	apiToken    string
	serviceName string

	Debug bool
}

func NewClient(baseUrl string, key string) *Client {
	baseURL, _ := url.Parse(baseUrl)

	c := &Client{
		Client:      http.DefaultClient,
		baseURL:     baseURL,
		apiKey:      key,
		apiToken:    viper.GetString("auth_key"),
		tenantId:    viper.GetString("tenant_id"),
		serviceName: "tenants",
		Debug:       false,
	}

	return c
}

func (c *Client) CreateAndDo(method, path string, data, options, resource interface{}) error {
	req, err := c.NewRequest(method, path, data)
	if err != nil {
		return err
	}

	err = c.Do(req, resource)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.serviceName, c.tenantId, path))
	if err != nil {
		return nil, err
	}

	// Make the full url based on the relative path
	u := c.baseURL.ResolveReference(rel)

	// Add custom options

	// A bit of JSON ceremony
	var js []byte = nil

	if body != nil {
		js, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(js))
	if err != nil {
		return nil, err
	}

	req.Header.Add("ONQLAVE-API-KEY", c.apiKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	if c.apiToken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = CheckResponseError(resp)
	if err != nil {
		return err
	}

	if v != nil {
		err := json.NewDecoder(resp.Body).Decode(&v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) Get(path string, resource, options interface{}) error {
	return c.CreateAndDo(http.MethodGet, path, nil, options, resource)
}

func (c *Client) Post(path string, data, resource, options interface{}) error {
	return c.CreateAndDo(http.MethodPost, path, data, options, resource)
}

func (c *Client) Put(path string, data, resource interface{}) error {
	return c.CreateAndDo(http.MethodPut, path, data, nil, resource)
}

func (c *Client) Delete(path string, options interface{}) error {
	return c.CreateAndDo(http.MethodDelete, path, nil, options, nil)
}
