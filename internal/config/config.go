package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	JWKS string `yaml:"jwks" env:"JWKS,secret"`
	PORT string `yaml:"port" env:"PORT,secret"`
}

// Load the configuration from a file.
func Load(file string) (*Config, error) {

	c := Config{}
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}
	return &c, err
}
