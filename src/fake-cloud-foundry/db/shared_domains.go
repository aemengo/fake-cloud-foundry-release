package db

import (
	"github.com/satori/go.uuid"
	"time"
)

type SharedDomain struct {
	Guid      string
	Name      string
	Internal  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (db *DB) loadSharedDomains() {
	now := time.Now().UTC()

	db.sharedDomains = []SharedDomain{{
		Guid:      uuid.NewV4().String(),
		Name:      db.config.Domain(),
		Internal:  false,
		CreatedAt: now,
		UpdatedAt: now,
	}}
}

func (db *DB) GetSharedDomains() []SharedDomain {
	return db.sharedDomains
}

func (db *DB) GetSharedDomainByGuid(guid string) (SharedDomain, bool) {
	for _, sharedDomain := range db.sharedDomains {
		if sharedDomain.Guid == guid {
			return sharedDomain, true
		}
	}

	return SharedDomain{}, false
}
