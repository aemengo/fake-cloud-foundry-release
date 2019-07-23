package uaa

import cfg "github.com/aemengo/fake-cloud-foundry/config"

type UAA struct {
	config cfg.Config
}

func New(config cfg.Config) *UAA {
	return &UAA{
		config: config,
	}
}

