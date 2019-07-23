package api

import cfg "github.com/aemengo/fake-cloud-foundry/config"

type API struct {
	config cfg.Config
}

func New(config cfg.Config) *API {
	return &API{
		config: config,
	}
}
