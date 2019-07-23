package db

import "github.com/satori/go.uuid"

type Org struct {
	Guid                        string
	Name                        string
	BillingEnabled              bool
	QuotaDefinitionGuid         string
	Status                      string
	DefaultIsolationSegmentGuid *string
}

func (db *DB) loadOrgs() {
	for _, org := range db.config.Orgs {
		db.orgs = append(db.orgs, Org{
			Guid:                        uuid.NewV4().String(),
			Name:                        org.Name,
			BillingEnabled:              false,
			QuotaDefinitionGuid:         "",
			Status:                      "active",
			DefaultIsolationSegmentGuid: nil,
		})
	}
}

func (db *DB) GetOrgs() []Org {
	return db.orgs
}

func (db *DB) GetOrgByName(name string) (Org, bool) {
	for _, org := range db.orgs {
		if org.Name == name {
			return org, true
		}
	}

	return Org{}, false
}

func (db *DB) GetOrgByGuid(guid string) (Org, bool) {
	for _, org := range db.orgs {
		if org.Guid == guid {
			return org, true
		}
	}

	return Org{}, false
}