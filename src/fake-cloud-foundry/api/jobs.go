package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func (a *API) Jobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")

	job, ok := a.database.GetJobByGuid(guid)
	if !ok {
		// do something
		return
	}

	resource := Resource{
		Metadata: Metadata{
			Guid:      job.Guid,
			URL:       fmt.Sprintf("/v2/jobs/%s", job.Guid),
			CreatedAt: job.CreatedAt.Format(time.RFC3339),
		},
		Entity: presentJob(job),
	}

	json.NewEncoder(w).Encode(resource)
}
