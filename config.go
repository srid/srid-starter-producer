package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port     string `envconfig:"PORT"`
	Username string `envconfig:"USERNAME"`
	Password string `envconfig:"PASSWORD"`
}

var config Config

func (c Config) validate() error {
	if c.Port == "" {
		return fmt.Errorf("$PORT is empty")
	}
	if c.Username == "" {
		return fmt.Errorf("$USERNAME is empty")
	}
	if c.Password == "" {
		return fmt.Errorf("$PASSWORD is empty")
	}
	return nil
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
