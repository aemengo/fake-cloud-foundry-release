package api

import (
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/aemengo/fake-cloud-foundry/db"
)

type API struct {
	config cfg.Config
	database     *db.DB
}

func New(config cfg.Config, database *db.DB) *API {
	return &API{
		config: config,
		database: database,
	}
}