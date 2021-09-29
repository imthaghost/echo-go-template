package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	localConfigFilePath = "/src/github.com/imthaghost/asus/config/%s.json"
	globalConfigFilePath =  "/etc/asus/config/%s.json"

	unspecifiedEnvironment = ""
	devEnvironment = "dev"
)

// Config ...
type Config struct {}

// GetEnvironment returns environment type
func GetEnvironment() string {
	env := os.Getenv("ENVIRONMENT")
	if env == unspecifiedEnvironment {
		return devEnvironment
	}
	return env
}

// LoadConfigForEnvironment pulls the Config data from the JSON config file
func LoadConfigForEnvironment(environment string) (*Config, error) {
	configFilePath, err := getConfigFilePath(environment)
	if err != nil {
		return nil, err
	}

	return loadConfig(configFilePath)
}

func getConfigFilePath(environment string) (string, error) {
	localFname := os.Getenv("GOPATH") + fmt.Sprintf(localConfigFilePath, environment)
	if _, err := os.Stat(localFname); !os.IsNotExist(err) {
		return localFname, nil
	}
	globalFname := fmt.Sprintf(globalConfigFilePath, environment)
	if _, err := os.Stat(globalFname); os.IsNotExist(err) {
		return "", err
	}
	return globalFname, nil
}


func loadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open config file: %v", err)
	}
	decoder := json.NewDecoder(file)
	configuration := &Config{}
	if err := decoder.Decode(configuration); err != nil {
		return nil, fmt.Errorf("Failed to parse configuration file: %v", err)
	}


	return configuration, nil
}
