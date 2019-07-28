package api

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

// Sample response
// {
//  "next_url": null,
//  "prev_url": null,
//  "resources": [
//    {
//      "entity": {
//        "internal": false,
//        "name": "dev.cfdev.sh",
//        "router_group_guid": null,
//        "router_group_type": null
//      },
//      "metadata": {
//        "created_at": "2019-07-22T01:55:53Z",
//        "guid": "e737e4af-3c9d-4aa5-b692-acc3c5fd81e1",
//        "updated_at": "2019-07-22T01:55:53Z",
//        "url": "/v2/shared_domains/e737e4af-3c9d-4aa5-b692-acc3c5fd81e1"
//      }
//    }
//  ],
//  "total_pages": 1,
//  "total_results": 1
//}

func (a *API) SharedDomains(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var resources []Resource

	for _, sharedDomain := range a.database.GetSharedDomains() {
		resources = append(resources, Resource{
			Metadata: Metadata{
				Guid:      sharedDomain.Guid,
				URL:       fmt.Sprintf("/v2/shared_domains/%s", sharedDomain.Guid),
				CreatedAt: sharedDomain.CreatedAt.Format(time.RFC3339),
				UpdatedAt: sharedDomain.UpdatedAt.Format(time.RFC3339),
			},
			Entity: presentSharedDomain(sharedDomain),
		})
	}

	response := newResponse(resources)
	json.NewEncoder(w).Encode(response)
}

func (a *API) SharedDomain(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")

	sharedDomain, ok := a.database.GetSharedDomainByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	resource := Resource{
		Metadata: Metadata{
			Guid:      sharedDomain.Guid,
			URL:       fmt.Sprintf("/v2/shared_domains/%s", sharedDomain.Guid),
			CreatedAt: sharedDomain.CreatedAt.Format(time.RFC3339),
			UpdatedAt: sharedDomain.UpdatedAt.Format(time.RFC3339),
		},
		Entity: presentSharedDomain(sharedDomain),
	}

	json.NewEncoder(w).Encode(resource)
}

func presentSharedDomain(sharedDomain db.SharedDomain) map[string]interface{} {
	return map[string]interface{}{
		"internal":          sharedDomain.Internal,
		"name":              sharedDomain.Name,
		"router_group_guid": nil,
		"router_group_type": nil,
	}
}
