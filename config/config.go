package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

//Config All config wrap
type Config struct {
	FeedURL      string             `yaml:"feedURL"`
	Transmission TransmissionConfig `yaml:"transmission"`
	SourceFolder string             `yaml:"sourceFolder"`
	TargetFolder string             `yaml:"targetFolder"`
	VideoExts    []string           `yaml:"videoExts"`
}

//TransmissionConfig Config about transmission
type TransmissionConfig struct {
	Host            string
	User            string
	Password        string
	RemoveCompletes bool `yaml:"removeCompletes"`
}

//Data is var with all configs
var Data Config

func init() {
	log.Println("initializing config...")
	dat, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Panic("error reading config file", err)
	}
	errorM := yaml.Unmarshal(dat, &Data)
	if errorM != nil {
		log.Panic("error parsing config file", errorM)
	}
	log.Printf("Config is %+v\n", Data)
}
