package db

import "github.com/satori/go.uuid"

type Space struct {
	Guid                     string
	Name                     string
	OrgGuid                  string
	SpaceQuotaDefinitionGuid *string
	IsolationSegmentGuid     *string
	AllowSSH                 bool
}

func (db *DB) loadSpaces() {
	for _, space := range db.config.Spaces {
		org, _ := db.GetOrgByName(space.Org)

		db.spaces = append(db.spaces, Space{
			Guid:                     uuid.NewV4().String(),
			Name:                     space.Name,
			OrgGuid:                  org.Guid,
			SpaceQuotaDefinitionGuid: nil,
			IsolationSegmentGuid:     nil,
			AllowSSH:                 false,
		})
	}
}

func (db *DB) GetSpacesByOrgGuid(orgGuid string) []Space {
	var spaces []Space

	for _, space := range db.spaces {
		if space.OrgGuid == orgGuid {
			spaces = append(spaces, space)
		}
	}

	return spaces
}