package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func (a *API) OrgSpaces(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		resources []Resource
		orgGuid   = ps.ByName("guid")
		spaces    = a.database.GetSpacesByOrgGuid(orgGuid)
	)

	for _, space := range spaces {
		resources = append(resources, Resource{
			Metadata: Metadata{
				Guid: space.Guid,
				URL: fmt.Sprintf("/v2/spaces/%s", space.Guid),
				CreatedAt: space.CreatedAt.Format(time.RFC3339),
				UpdatedAt: space.UpdatedAt.Format(time.RFC3339),
			},
			Entity: presentSpace(space),
		})
	}

	response := newResponse(resources)
	json.NewEncoder(w).Encode(response)
}

func presentSpace(space db.Space) map[string]interface{} {
	return map[string]interface{}{
		"name":                        space.Name,
		"organization_guid":           space.OrgGuid,
		"space_quota_definition_guid": space.SpaceQuotaDefinitionGuid,
		"isolation_segment_guid":      space.IsolationSegmentGuid,
		"allow_ssh":                   space.AllowSSH,
		"organization_url":            fmt.Sprintf("/v2/organizations/%s", space.OrgGuid),
		"developers_url":              fmt.Sprintf("/v2/spaces/%s/developers", space.Guid),
		"managers_url":                fmt.Sprintf("/v2/spaces/%s/managers", space.Guid),
		"auditors_url":                fmt.Sprintf("/v2/spaces/%s/auditors", space.Guid),
		"apps_url":                    fmt.Sprintf("/v2/spaces/%s/apps", space.Guid),
		"routes_url":                  fmt.Sprintf("/v2/spaces/%s/routes", space.Guid),
		"domains_url":                 fmt.Sprintf("/v2/spaces/%s/domains", space.Guid),
		"service_instances_url":       fmt.Sprintf("/v2/spaces/%s/service_instances", space.Guid),
		"app_events_url":              fmt.Sprintf("/v2/spaces/%s/app_events", space.Guid),
		"events_url":                  fmt.Sprintf("/v2/spaces/%s/events", space.Guid),
		"security_groups_url":         fmt.Sprintf("/v2/spaces/%s/security_groups", space.Guid),
		"staging_security_groups_url": fmt.Sprintf("/v2/spaces/%s/staging_security_groups", space.Guid),
	}
}
