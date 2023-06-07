package api

import (
	"fmt"

	"github.com/spf13/viper"
)

type ServiceName string

const (
	BillingName      ServiceName = "billings"
	TenantName       ServiceName = "tenants"
	RegistrationName ServiceName = "registrations"
)

func (sn ServiceName) String() string {
	return string(sn)
}

func IsLocal() bool {
	return viper.GetString("env") == "local"
}

func IsDev() bool {
	return viper.GetString("env") == "dev"
}

func IsProd() bool {
	return viper.GetString("env") == "prod"
}

func GetBaseUrl(serviceName string) string {
	const defaultURL = "http://localhost:8081"
	apiBaseUrl := viper.GetStringMapString("api_base_url")
	if url, ok := apiBaseUrl[serviceName]; ok {
		return url
	}
	return defaultURL
}
func UrlBuilder(serviceName string) string {
	if IsLocal() {
		return fmt.Sprintf("%s/%s", GetBaseUrl(serviceName), serviceName)
	}
	return fmt.Sprintf("%s/%s", viper.GetString("api_base_url"), serviceName)
}
