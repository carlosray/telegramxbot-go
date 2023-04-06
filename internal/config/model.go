package config

import (
	"telegramxbot/internal/model"
)

type Config struct {
	Bot struct {
		Debug    bool
		Username string
		Token    string
	}

	Handlers []struct {
		Name       string
		Properties map[string]interface{}
	}

	model.HandlePolicy `yaml:"handle_policy"`

	UpdateConfig struct {
		Timeout int
	} `yaml:"update_config"`

	Log struct {
		Level string
	}
}
