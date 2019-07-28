package api

import (
	"encoding/json"
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//{
//  "resources": [
//    {
//      "details": null,
//      "disk_quota": 1073741824,
//      "fds_quota": 16384,
//      "host": "10.144.0.13",
//      "index": 0,
//      "instance_ports": [
//        {
//          "external": 61000,
//          "internal": 8080
//        },
//        {
//          "external": 61001,
//          "internal": 2222
//        }
//      ],
//      "isolation_segment": null,
//      "mem_quota": 268435456,
//      "state": "RUNNING",
//      "type": "web",
//      "uptime": 1,
//      "usage": {
//        "cpu": 0.0,
//        "disk": 94195712,
//        "mem": 106496,
//        "time": "2019-07-22T20:47:44+00:00"
//      }
//    }
//  ]
//}

func (a *API) V3ProcessStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := newV3Response(
		[]V3Resource{presentProcessStat(a.config)},
		r.URL.String(),
	)

	json.NewEncoder(w).Encode(response)
}

func presentProcessStat(config cfg.Config) map[string]interface{} {
	return map[string]interface{}{
		"details":           nil,
		"disk_quota":        0,
		"fds_quota":         0,
		"host":              "10.0.0.1", //TODO <-- real value
		"index":             0,
		"instance_ports":    []interface{}{},
		"isolation_segment": nil,
		"mem_quota":         0,
		"state":             "RUNNING",
		"type":              "web",
		"uptime":            1,
		"usage": map[string]interface{}{
			"cpu":  0.0,
			"disk": 0,
			"mem":  0,
			"time": time.Now().UTC().Format(time.RFC3339),
		},
	}
}
