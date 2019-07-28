package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Users []struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	} `json:"users"`
	Orgs []struct {
		Name string `json:"name"`
	} `json:"orgs"`
	Spaces []struct {
		Name string `json:"name"`
		Org  string `json:"org"`
	}
}

func New(path string) (Config, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) Domain() string {
	return c.Host + ":" + c.Port
}
