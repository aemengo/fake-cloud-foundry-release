package api

import (
	"encoding/json"
	"fmt"
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//{
//  "pagination": {
//    "first": {
//      "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/processes?page=1&per_page=50"
//    },
//    "last": {
//      "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/processes?page=1&per_page=50"
//    },
//    "next": null,
//    "previous": null,
//    "total_pages": 1,
//    "total_results": 1
//  },
//  "resources": [
//    {
//      "command": "[PRIVATE DATA HIDDEN IN LISTS]",
//      "created_at": "2019-07-22T20:46:47Z",
//      "disk_in_mb": 1024,
//      "guid": "cb05ab74-3a08-453b-ab14-62d0274c3244",
//      "health_check": {
//        "data": {
//          "invocation_timeout": null,
//          "timeout": null
//        },
//        "type": "port"
//      },
//      "instances": 1,
//      "links": {
//        "app": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244"
//        },
//        "scale": {
//          "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244/actions/scale",
//          "method": "POST"
//        },
//        "self": {
//          "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244"
//        },
//        "space": {
//          "href": "https://api.dev.cfdev.sh/v3/spaces/2f732a38-8035-4696-a243-c52368c2b190"
//        },
//        "stats": {
//          "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244/stats"
//        }
//      },
//      "memory_in_mb": 256,
//      "metadata": {
//        "annotations": {},
//        "labels": {}
//      },
//      "relationships": {
//        "revision": null
//      },
//      "type": "web",
//      "updated_at": "2019-07-22T20:46:51Z"
//    }
//  ]
//}

func (a *API) V3AppProcesses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")

	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	response := newV3Response(
		[]V3Resource{presentV3AppProcess(app, a.config)},
		r.URL.String(),
	)

	json.NewEncoder(w).Encode(response)
}

//{
//  "command": "bundle exec rackup config.ru -p $PORT",
//  "created_at": "2019-07-22T20:46:47Z",
//  "disk_in_mb": 1024,
//  "guid": "cb05ab74-3a08-453b-ab14-62d0274c3244",
//  "health_check": {
//    "data": {
//      "invocation_timeout": null,
//      "timeout": null
//    },
//    "type": "port"
//  },
//  "instances": 1,
//  "links": {
//    "app": {
//      "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244"
//    },
//    "scale": {
//      "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244/actions/scale",
//      "method": "POST"
//    },
//    "self": {
//      "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244"
//    },
//    "space": {
//      "href": "https://api.dev.cfdev.sh/v3/spaces/2f732a38-8035-4696-a243-c52368c2b190"
//    },
//    "stats": {
//      "href": "https://api.dev.cfdev.sh/v3/processes/cb05ab74-3a08-453b-ab14-62d0274c3244/stats"
//    }
//  },
//  "memory_in_mb": 256,
//  "metadata": {
//    "annotations": {},
//    "labels": {}
//  },
//  "relationships": {
//    "revision": null
//  },
//  "type": "web",
//  "updated_at": "2019-07-22T20:46:51Z"
//}

func (a *API) V3AppProcess(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guid := ps.ByName("guid")
	//process := ps.ByName("process")

	app, ok := a.database.GetAppByGuid(guid)
	if !ok {
		//TODO something
		return
	}

	response := presentV3AppProcess(app, a.config)
	json.NewEncoder(w).Encode(response)
}


func presentV3AppProcess(app db.App, config cfg.Config) map[string]interface{} {
	return map[string]interface{}{
		"command":    app.DetectedStartCommand,
		"created_at": app.CreatedAt.Format(time.RFC3339),
		"disk_in_mb": app.DiskQuota,
		"guid":       app.Guid,
		"health_check": map[string]interface{}{
			"data": map[string]interface{}{
				"invocation_timeout": app.HealthCheckTimeout,
				"timeout":            app.HealthCheckTimeout,
			},
			"type": app.HealthCheckType,
		},
		"instances": app.Instances,
		"links": map[string]interface{}{
			"app": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s", config.Domain(), app.Guid),
			},
			"scale": map[string]interface{}{
				"href":   fmt.Sprintf("http://%s/v3/processes/%s/actions/scale", config.Domain(), app.Guid),
				"method": http.MethodPost,
			},
			"self": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/processes/%s", config.Domain(), app.Guid),
			},
			"space": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/spaces/%s", config.Domain(), app.SpaceGuid),
			},
			"stats": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/processes/%s/stats", config.Domain(), app.Guid),
			},
		},
		"memory_in_mb": app.Memory,
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{},
			"labels":      map[string]interface{}{},
		},
		"relationships": map[string]interface{}{
			"revision": nil,
		},
		"type":       "web",
		"updated_at": app.UpdatedAt.Format(time.RFC3339),
	}
}