package config

import "log"

// SetupParams ...
type SetupParams struct {
	Config *Config
}

// SetupAsus configuration values
func SetupAsus() *SetupParams {
	env := GetEnvironment()

	if env == "production" {
		env = "prod"
	}
	cfg, err := LoadConfigForEnvironment(env)
	if err != nil {
		log.Fatal(err)
	}

	return &SetupParams{Config: cfg}
}