// Package config provides a configuration file for the application
package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

// Config is the configuration file for the application
type Config struct {
	Repository string
}

// FileExists checks if the config file exists
func FileExists() bool {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	configPath := path.Join(homedir, ".config", "mattd", "config.yml")
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// FromFile reads the config file and returns a Config struct
func FromFile(path string) (*Config, error) {
	cfgfile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(cfgfile, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

// ToFile writes the Config struct to the config file
func ToFile(filepath string, cfg *Config) error {
	bytes, err := yaml.Marshal(cfg)

	if err != nil {
		return err
	}

	dir := path.Dir(filepath)

	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filepath, bytes, 0644); err != nil {
		return err
	}

	return nil
}
