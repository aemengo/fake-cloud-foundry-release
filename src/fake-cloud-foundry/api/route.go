package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//Sample response
//{
//  "next_url": null,
//  "prev_url": null,
//  "resources": null,
//  "total_pages": 1,
//  "total_results": 0
//}

func (a *API) GetRoutes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var resources []Resource

	for _, _ = range a.database.GetRoutes() {
		resources = append(resources, Resource{})
	}

	response := newResponse(resources)
	json.NewEncoder(w).Encode(response)
}

type postRoutesForm struct {
	Host       string  `json:"host"`
	Port       *string `json:"port"`
	SpaceGuid  string  `json:"space_guid"`
	DomainGuid string  `json:"domain_guid"`
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

//TODO: add routes to router
func (a *API) PostRoutes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var f postRoutesForm
	json.NewDecoder(r.Body).Decode(&f)

	route := a.database.CreateRoute(db.Route{
		DomainGuid:          f.DomainGuid,
		Host:                f.Host,
		Path:                "",
		Port:                f.Port,
		ServiceInstanceGuid: nil,
		SpaceGuid:           f.SpaceGuid,
	})

	now := time.Now().Format(time.RFC3339)
	resource := Resource{
		Metadata: Metadata{
			Guid:      route.Guid,
			URL:       fmt.Sprintf("/v2/routes/%s", route.Guid),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Entity: presentRoute(route),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

func presentRoute(route db.Route) map[string]interface{} {
	return map[string]interface{}{
		"apps_url":              fmt.Sprintf("/v2/routes/%s/apps", route.Guid),
		"domain_guid":           route.DomainGuid,
		"domain_url":            fmt.Sprintf("/v2/shared_domains/%s", route.DomainGuid),
		"host":                  route.Host,
		"path":                  route.Path,
		"port":                  route.Port,
		"route_mappings_url":    fmt.Sprintf("/v2/routes/%s/route_mappings", route.Guid),
		"service_instance_guid": route.ServiceInstanceGuid,
		"space_guid":            route.SpaceGuid,
		"space_url":             fmt.Sprintf("/v2/spaces/%s", route.SpaceGuid),
	}
}
