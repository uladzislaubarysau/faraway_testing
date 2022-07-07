package main

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host string `envconfig:"HOST"`
	Port string `envconfig:"PORT"`
}

func loadConfig() (config, error) {
	var conf config

	err := envconfig.Process("WOW", &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
