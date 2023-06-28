package common

import (
	"fmt"
	"os"
	"runtime"
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
	dir, _ := os.UserConfigDir()

	separator := "/"
	if runtime.GOOS == OSWindows {
		separator = "\\"
	}

	return fmt.Sprintf("%s%s%s%s", dir, separator, configDir, separator)
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
