package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port     string `envconfig:"PORT"`
	USERNAME string `envconfig:"USERNAME"`
	PASSWORD string `envconfig:"PASSWORD"`
}

var config Config

func (c config) validate() error {
	if c.Port == "" {
		return fmt.Errorf("$PORT is empty")
	}
	if c.Username == "" {
		return fmt.Errorf("$USERNAME is empty")
	}
	if c.Password == "" {
		return fmt.Errorf("$PASSWORD is empty")
	}
}

func init() {
	err := envconfig.Process("producer", &config)
	if err == nil {
		err = config.validate()

	}
	if err != nil {
		log.Fatal(err.Error())
	}
}
