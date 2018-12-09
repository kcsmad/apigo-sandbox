package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Server string
	Database string
}

func (config *Config) ReadConfigFile() {
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatal(err)
	}
}