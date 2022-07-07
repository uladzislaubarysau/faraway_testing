package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ServerHost   string        `envconfig:"SERVER_HOST"`
	ServerPort   string        `envconfig:"SERVER_PORT"`
	TimeForRetry time.Duration `envconfig:"TIME_FOR_RETRY"`
}

func loadConfig() (config, error) {
	var conf config

	err := envconfig.Process("WOW", &conf)

	if err != nil {
		return conf, err
	}

	return conf, nil
}
