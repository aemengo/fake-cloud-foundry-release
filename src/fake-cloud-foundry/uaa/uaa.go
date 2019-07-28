package uaa

import (
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/aemengo/fake-cloud-foundry/db"
)

type UAA struct {
	config   cfg.Config
	database *db.DB
}

func New(config cfg.Config, database *db.DB) *UAA {
	return &UAA{
		config:   config,
		database: database,
	}
}
