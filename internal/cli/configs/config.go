package configs

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type Config struct {
	GatewayAddress string
	Dev            bool
}

// FromFile returns a configuration object from a given file path and errors
// if the file does not exists of it the file contexts can't be unmarshaled.
// expected format is yaml.
func FromFile(path string) (*Config, error) {
	cfgfile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := yaml.Unmarshal(cfgfile, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// ToFile marshals a configuration object to the given filepath using the yaml
// format
func ToFile(filePath string, cfg *Config) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	d := path.Dir(filePath)
	if _, err := os.Stat(d); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}
	if err := os.WriteFile(filePath, b, 0644); err != nil {
		return err
	}
	return nil
}

func CreateFile(filePath string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		d := path.Dir(filePath)
		if _, err := os.Stat(d); errors.Is(err, os.ErrNotExist) {
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	err := viper.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
