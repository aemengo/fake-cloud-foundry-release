package db

import (
	cfg "github.com/aemengo/fake-cloud-foundry/config"
)

type DB struct {
	config        cfg.Config
	users         []User
	orgs          []Org
	spaces        []Space
	apps          []App
	sharedDomains []SharedDomain
	routes        []Route
	jobs          []Job
}

func New(config cfg.Config) *DB {
	db := &DB{
		config: config,
	}

	db.load()
	return db
}

func (db *DB) load() {
	db.loadUsers()
	db.loadOrgs()
	db.loadSpaces()
	db.loadSharedDomains()
}
