package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const (
	appConfigEnv = "APPLICATION_CONFIG"
	tokenEnv     = "BOT_TOKEN"
)

func LoadConfig() *Config {
	filename := getEnv(appConfigEnv, "configs/default.yaml")
	cfg, err := readConfig(filename)
	cfg.setUpSecrets()
	if err != nil {
		log.Panicln(err)
	}
	return cfg
}

func readConfig(filename string) (*Config, error) {
	cfg := &Config{}
	file, err := os.ReadFile(filename)
	if err == nil {
		log.Printf("Read config: \n%s\n", file)
		err = yaml.Unmarshal(file, cfg)
	}
	return cfg, err
}

func (c *Config) setUpSecrets() {
	c.Bot.Token = getEnv(tokenEnv, c.Bot.Token)
}
