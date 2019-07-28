package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//{
//"total_results": 2,
//"total_pages": 1,
//"prev_url": null,
//"next_url": null,
//"resources": [
//  {
//    "metadata": {
//      "guid": "157b05e0-d2a9-47f1-9794-cbbb6c8ae484",
//      "url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484",
//      "created_at": "2019-07-22T01:59:26Z",
//      "updated_at": "2019-07-22T01:59:26Z"
//    },
//    "entity": {
//      "name": "cfdev-org",
//      "billing_enabled": false,
//      "quota_definition_guid": "2c18e346-55f7-49a6-856d-f0a1d2d8418f",
//      "status": "active",
//      "default_isolation_segment_guid": null,
//      "quota_definition_url": "/v2/quota_definitions/2c18e346-55f7-49a6-856d-f0a1d2d8418f",
//      "spaces_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/spaces",
//      "domains_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/domains",
//      "private_domains_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/private_domains",
//      "users_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/users",
//      "managers_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/managers",
//      "billing_managers_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/billing_managers",
//      "auditors_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/auditors",
//      "app_events_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/app_events",
//      "space_quota_definitions_url": "/v2/organizations/157b05e0-d2a9-47f1-9794-cbbb6c8ae484/space_quota_definitions"
//    }
//    },
//    {
//    "metadata": {
//      "guid": "a54044e4-7f2a-4942-a104-25a64243c572",
//      "url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572",
//      "created_at": "2019-07-22T01:55:53Z",
//      "updated_at": "2019-07-22T01:55:53Z"
//    },
//    "entity": {
//      "name": "system",
//      "billing_enabled": false,
//      "quota_definition_guid": "2c18e346-55f7-49a6-856d-f0a1d2d8418f",
//      "status": "active",
//      "default_isolation_segment_guid": null,
//      "quota_definition_url": "/v2/quota_definitions/2c18e346-55f7-49a6-856d-f0a1d2d8418f",
//      "spaces_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/spaces",
//      "domains_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/domains",
//      "private_domains_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/private_domains",
//      "users_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/users",
//      "managers_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/managers",
//      "billing_managers_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/billing_managers",
//      "auditors_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/auditors",
//      "app_events_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/app_events",
//      "space_quota_definitions_url": "/v2/organizations/a54044e4-7f2a-4942-a104-25a64243c572/space_quota_definitions"
//    }
//  }
//]
//}

func (a *API) Orgs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var resources []Resource

	for _, org := range a.database.GetOrgs() {
		resources = append(resources, Resource{
			Metadata: Metadata{
				Guid:      org.Guid,
				URL:       fmt.Sprintf("/v2/organizations/%s", org.Guid),
				CreatedAt: org.CreatedAt.Format(time.RFC3339),
				UpdatedAt: org.UpdatedAt.Format(time.RFC3339),
			},
			Entity: presentOrg(org),
		})
	}

	response := newResponse(resources)
	json.NewEncoder(w).Encode(response)
}

func presentOrg(org db.Org) map[string]interface{} {
	return map[string]interface{}{
		"name":                           org.Name,
		"billing_enabled":                org.BillingEnabled,
		"quota_definition_guid":          org.QuotaDefinitionGuid,
		"status":                         org.Status,
		"default_isolation_segment_guid": org.DefaultIsolationSegmentGuid,
		"quota_definition_url":           fmt.Sprintf("/v2/quota_definitions/%s", ""),
		"spaces_url":                     fmt.Sprintf("/v2/organizations/%s/spaces", org.Guid),
		"domains_url":                    fmt.Sprintf("/v2/organizations/%s/domains", org.Guid),
		"private_domains_url":            fmt.Sprintf("/v2/organizations/%s/private_domains", org.Guid),
		"users_url":                      fmt.Sprintf("/v2/organizations/%s/users", org.Guid),
		"managers_url":                   fmt.Sprintf("/v2/organizations/%s/managers", org.Guid),
		"billing_managers_url":           fmt.Sprintf("/v2/organizations/%s/billing_managers", org.Guid),
		"auditors_url":                   fmt.Sprintf("/v2/organizations/%s/auditors", org.Guid),
		"app_events_url":                 fmt.Sprintf("/v2/organizations/%s/app_events", org.Guid),
		"space_quota_definitions_url":    fmt.Sprintf("/v2/organizations/%s/space_quota_definitions", org.Guid),
	}
}
