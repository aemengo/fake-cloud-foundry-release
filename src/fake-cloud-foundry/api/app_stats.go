package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//{
//  "0": {
//    "isolation_segment": null,
//    "state": "RUNNING",
//    "stats": {
//      "disk_quota": 1073741824,
//      "fds_quota": 16384,
//      "host": "10.144.0.13",
//      "mem_quota": 268435456,
//      "name": "dora",
//      "port": 61000,
//      "uptime": 1,
//      "uris": [
//        "dora.dev.cfdev.sh"
//      ],
//      "usage": {
//        "cpu": 0.0,
//        "disk": 94195712,
//        "mem": 106496,
//        "time": "2019-07-22T20:47:45+00:00"
//      }
//    }
//  }
//}

func (a *API) AppStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")

	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	info := map[string]interface{}{
		"0": map[string]interface{}{
			"isolation_segment": nil,
			"state":             app.State,
			"stats": map[string]interface{}{
				"disk_quota": 0,
				"fds_quota":  0,
				"host":       "10.0.0.1",
				"mem_quota":  0,
				"name":       app.Name,
				"port":       0,
				"uptime":     1,
				"uris":       []string{},
				"usage": map[string]interface{}{
					"cpu":  0.0,
					"disk": 0,
					"mem":  0,
					"time": time.Now().Format(time.RFC3339),
				},
			},
		},
	}

	json.NewEncoder(w).Encode(info)
}
