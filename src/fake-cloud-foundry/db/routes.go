package db

import "github.com/satori/go.uuid"

type Route struct {
	Guid                string
	DomainGuid          string
	DomainURL           string
	Host                string
	Path                string
	Port                *string
	ServiceInstanceGuid *string
	SpaceGuid           string
}

func (db *DB) GetRoutes() []Route {
	return db.routes
}

//{
//  "entity": {
//    "apps_url": "/v2/routes/a2666c77-d904-4888-96bf-14160edefb5e/apps",
//    "domain_guid": "e737e4af-3c9d-4aa5-b692-acc3c5fd81e1",
//    "domain_url": "/v2/shared_domains/e737e4af-3c9d-4aa5-b692-acc3c5fd81e1",
//    "host": "my-different-app-2",
//    "path": "",
//    "port": null,
//    "route_mappings_url": "/v2/routes/a2666c77-d904-4888-96bf-14160edefb5e/route_mappings",
//    "service_instance_guid": null,
//    "space_guid": "2f732a38-8035-4696-a243-c52368c2b190",
//    "space_url": "/v2/spaces/2f732a38-8035-4696-a243-c52368c2b190"
//  },
//  "metadata": {
//    "created_at": "2019-07-23T06:38:49Z",
//    "guid": "a2666c77-d904-4888-96bf-14160edefb5e",
//    "updated_at": "2019-07-23T06:38:49Z",
//    "url": "/v2/routes/a2666c77-d904-4888-96bf-14160edefb5e"
//  }
//}

func (db *DB) CreateRoute(route Route) Route {
	route.Guid = uuid.NewV4().String()

	db.routes = append(db.routes, route)
	return route
}
