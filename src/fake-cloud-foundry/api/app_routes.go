package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func (a *API) AppRoutes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")
	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	sharedDomain := a.database.GetSharedDomains()[0]

	resource := Resource{
		Metadata: Metadata{
			Guid:      app.Guid,
			URL:       fmt.Sprintf("/v2/routes/%s", app.Guid),
			CreatedAt: app.CreatedAt.Format(time.RFC3339),
			UpdatedAt: app.CreatedAt.Format(time.RFC3339),
		},
		Entity: presentAppRoute(app, sharedDomain.Guid),
	}

	response := newResponse([]Resource{resource})
	json.NewEncoder(w).Encode(response)
}

func presentAppRoute(app db.App, sharedDomainGuid string) map[string]interface{} {
	return map[string]interface{}{
		"apps_url":              "/v2/routes/25c7141c-414a-4abf-b7b3-991fe28350a4/apps",
		"domain_guid":           sharedDomainGuid,
		"domain_url":            fmt.Sprintf("/v2/shared_domains/%s", sharedDomainGuid),
		"host":                  app.Name,
		"path":                  "",
		"port":                  nil,
		"route_mappings_url":    "/v2/routes/25c7141c-414a-4abf-b7b3-991fe28350a4/route_mappings",
		"service_instance_guid": nil,
		"space_guid":            app.SpaceGuid,
		"space_url":             fmt.Sprintf("/v2/spaces/%s", app.SpaceGuid),
	}
}
