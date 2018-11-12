package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	FeedURL      string             `yaml:"feedURL"`
	Transmission TransmissionConfig `yaml:"transmission"`
}

type TransmissionConfig struct {
	Host            string
	User            string
	Password        string
	RemoveCompletes bool `yaml:"removeCompletes"`
}

var Data Config

func init() {
	log.Println("initializing config...")
	dat, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Panic("error rading config file", err)
	}
	errorM := yaml.Unmarshal(dat, &Data)
	if errorM != nil {
		log.Panic("error parsing config file", errorM)
	}
}
