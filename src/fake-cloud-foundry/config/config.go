package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Domain string `json:"domain"`
	Orgs   []struct {
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
