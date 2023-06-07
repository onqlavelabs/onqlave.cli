package common

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

var (
	BaseUrlLocal = map[string]string{
		"tenants":       "http://localhost:8083",
		"billings":      "http://localhost:8084",
		"registrations": "http://localhost:8081",
	}
)

func GetConfigFilePath() string {
	return GetConfigDir() + ConfigFile
}

func GetConfigDir() string {
	dir, _ := os.UserHomeDir()
	return dir + configDir
}

func IsEnvironmentConfigured() bool {
	return viper.Get(FlagApiBaseUrl) != nil
}

func IsLoggedIn() bool {
	return viper.GetString(FlagAuthKey) != "" && viper.Get(FlagApiBaseUrl) != nil
}

func IsSupportedEnv() bool {
	env := os.Getenv(strings.ToUpper(FlagEnv))
	return env != "" && env != EnvDev && env != EnvProd && env != EnvLocal
}
