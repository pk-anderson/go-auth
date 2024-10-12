package config

import (
	"io/ioutil"
	"log"

	"github.com/pk-anderson/go-auth/interfaces"
	"gopkg.in/yaml.v2"
)

func LoadConfig(filename string) interfaces.Config {
	cfg := interfaces.Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Error parsing YAMl file: %v", err)
	}

	return cfg
}
