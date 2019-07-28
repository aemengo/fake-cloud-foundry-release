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

//Sample response
//{
//  "next_url": null,
//  "prev_url": null,
//  "resources": null,
//  "total_pages": 1,
//  "total_results": 0
//}

func (a *API) GetApps(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var resources []Resource

	for _, _ = range a.database.GetApps() {
		resources = append(resources, Resource{})
	}

	response := newResponse(resources)
	json.NewEncoder(w).Encode(response)
}

//RESPONSE: [2019-07-28T19:18:46-04:00]
//HTTP/1.1 201 Created
//Content-Length: 1688
//Content-Type: application/json;charset=utf-8
//Date: Mon, 22 Jul 2019 20:46:47 GMT
//Location: /v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244
//Server: nginx
//X-Content-Type-Options: nosniff
//X-Vcap-Request-Id: b9ea1991-9c5e-45f4-56b3-31fbe5ed4731::14b2f146-dfb2-48da-a489-ddf2afbe7af9
//{
//  "entity": {
//    "buildpack": null,
//    "command": null,
//    "console": false,
//    "debug": null,
//    "detected_buildpack": null,
//    "detected_buildpack_guid": null,
//    "detected_start_command": "",
//    "diego": true,
//    "disk_quota": 1024,
//    "docker_credentials": {
//      "password": "[PRIVATE DATA HIDDEN]",
//      "username": null
//    },
//    "docker_image": null,
//    "enable_ssh": true,
//    "environment_json": {},
//    "events_url": "/v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/events",
//    "health_check_http_endpoint": null,
//    "health_check_timeout": null,
//    "health_check_type": "port",
//    "instances": 1,
//    "memory": 256,
//    "name": "dora",
//    "package_state": "PENDING",
//    "package_updated_at": null,
//    "ports": [
//      8080
//    ],
//    "production": false,
//    "route_mappings_url": "/v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/route_mappings",
//    "routes_url": "/v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/routes",
//    "service_bindings_url": "/v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/service_bindings",
//    "space_guid": "2f732a38-8035-4696-a243-c52368c2b190",
//    "space_url": "/v2/spaces/2f732a38-8035-4696-a243-c52368c2b190",
//    "stack_guid": "91cff036-4ae6-4f52-b575-3cf5f77889c6",
//    "stack_url": "/v2/stacks/91cff036-4ae6-4f52-b575-3cf5f77889c6",
//    "staging_failed_description": null,
//    "staging_failed_reason": null,
//    "staging_task_id": null,
//    "state": "STOPPED",
//    "version": "0fbce5e5-3058-4e66-a67f-b217d709d357"
//  },
//  "metadata": {
//    "created_at": "2019-07-22T20:46:47Z",
//    "guid": "cb05ab74-3a08-453b-ab14-62d0274c3244",
//    "updated_at": "2019-07-22T20:46:47Z",
//    "url": "/v2/apps/cb05ab74-3a08-453b-ab14-62d0274c3244"
//  }
//}

type postAppsForm struct {
	Name      string `json:"name"`
	SpaceGuid string `json:"space_guid"`
}

func (a *API) PostApps(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	defer r.Body.Close()

	var f postAppsForm
	json.NewDecoder(r.Body).Decode(&f)

	app := a.database.CreateApp(db.App{
		Name:      f.Name,
		SpaceGuid: f.SpaceGuid,
	})

	now := time.Now().Format(time.RFC3339)
	resource := Resource{
		Metadata: Metadata{
			Guid:      app.Guid,
			URL:       fmt.Sprintf("/v2/apps/%s", app.Guid),
			CreatedAt: now,
			UpdatedAt: now,
		},
		Entity: presentApp(app),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resource)
}

//{
//  "pagination": {
//    "first": {
//      "href": "https://api.dev.cfdev.sh/v3/apps?names=dora&page=1&per_page=50&space_guids=2f732a38-8035-4696-a243-c52368c2b190"
//    },
//    "last": {
//      "href": "https://api.dev.cfdev.sh/v3/apps?names=dora&page=1&per_page=50&space_guids=2f732a38-8035-4696-a243-c52368c2b190"
//    },
//    "next": null,
//    "previous": null,
//    "total_pages": 1,
//    "total_results": 1
//  },
//  "resources": [
//    {
//      "created_at": "2019-07-22T20:46:47Z",
//      "guid": "cb05ab74-3a08-453b-ab14-62d0274c3244",
//      "lifecycle": {
//        "data": {
//          "buildpacks": null,
//          "stack": "cflinuxfs3"
//        },
//        "type": "buildpack"
//      },
//      "links": {
//        "current_droplet": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/droplets/current"
//        },
//        "deployed_revisions": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/revisions/deployed"
//        },
//        "droplets": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/droplets"
//        },
//        "environment_variables": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/environment_variables"
//        },
//        "packages": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/packages"
//        },
//        "processes": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/processes"
//        },
//        "revisions": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/revisions"
//        },
//        "route_mappings": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/route_mappings"
//        },
//        "self": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244"
//        },
//        "space": {
//          "href": "https://api.dev.cfdev.sh/v3/spaces/2f732a38-8035-4696-a243-c52368c2b190"
//        },
//        "start": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/actions/start",
//          "method": "POST"
//        },
//        "stop": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/actions/stop",
//          "method": "POST"
//        },
//        "tasks": {
//          "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/tasks"
//        }
//      },
//      "metadata": {
//        "annotations": {},
//        "labels": {}
//      },
//      "name": "dora",
//      "relationships": {
//        "space": {
//          "data": {
//            "guid": "2f732a38-8035-4696-a243-c52368c2b190"
//          }
//        }
//      },
//      "state": "STARTED",
//      "updated_at": "2019-07-22T20:47:35Z"
//    }
//  ]
//}

func (a *API) V3Apps(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.URL.Query().Get("names")
	spaceGuid := r.URL.Query().Get("space_guids")
	
	app, ok := a.database.GetAppByNameAndSpaceGuid(name, spaceGuid)
	if !ok {
		//TODO something
		return
	}

	response := newV3Response(
		[]V3Resource{presentV3App(app, a.config)},
		r.URL.String(),
	)

	json.NewEncoder(w).Encode(response)
}

func presentV3App(app db.App, config cfg.Config) map[string]interface{} {
	return map[string]interface{}{
		"created_at": app.CreatedAt.Format(time.RFC3339),
		"guid": app.Guid,
		"lifecycle": map[string]interface{}{
			"data":  map[string]interface{}{
				"buildpacks": nil,
				"stack": "cflinuxfs3",
			},
			"type": "buildpack",
		},
		"links":  map[string]interface{}{
			"current_droplet": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/droplets/current", config.Domain(), app.Guid),
			},
			"deployed_revisions":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/revisions/deployed", config.Domain(), app.Guid),
			},
			"droplets":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/droplets", config.Domain(), app.Guid),
			},
			"environment_variables":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/environment_variables", config.Domain(), app.Guid),
			},
			"packages":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/packages", config.Domain(), app.Guid),
			},
			"processes":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/processes", config.Domain(), app.Guid),
			},
			"revisions":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/revisions", config.Domain(), app.Guid),
			},
			"route_mappings":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/route_mappings", config.Domain(), app.Guid),
			},
			"self":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s", config.Domain(), app.Guid),
			},
			"space":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/spaces/%s", config.Domain(), app.SpaceGuid),
			},
			"start":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/actions/start", config.Domain(), app.Guid),
				"method": http.MethodPost,
			},
			"stop":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/actions/stop", config.Domain(), app.Guid),
				"method": http.MethodPost,
			},
			"tasks":  map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/%s/tasks", config.Domain(), app.Guid),
			},
		},
		"metadata": map[string]interface{}{
			"annotations": map[string]interface{}{},
			"labels": map[string]interface{}{},
		},
		"name": app.Name,
		"relationships": map[string]interface{}{
			"space": map[string]interface{}{
				"data": map[string]interface{}{
					"guid": app.SpaceGuid,
				},
			},
		},
		"state": app.State,
		"updated_at": app.UpdatedAt.Format(time.RFC3339),
	}
}

func presentApp(app db.App) map[string]interface{} {
	return map[string]interface{}{
		"buildpack":               app.Buildpack,
		"command":                 app.Command,
		"console":                 app.Console,
		"debug":                   app.Debug,
		"detected_buildpack":      app.DetectedBuildpack,
		"detected_buildpack_guid": app.DetectedBuildpackGuid,
		"detected_start_command":  app.DetectedStartCommand,
		"diego":                   true,
		"disk_quota":              app.DiskQuota,
		"docker_credentials": map[string]interface{}{
			"password": "",
			"username": nil,
		},
		"docker_image":               nil,
		"enable_ssh":                 app.EnableSSH,
		"environment_json":           app.EnvironmentJSON,
		"events_url":                 fmt.Sprintf("/v2/apps/%s/events", app.Guid),
		"health_check_http_endpoint": app.HealthCheckHTTPEndpoint,
		"health_check_timeout":       app.HealthCheckTimeout,
		"health_check_type":          app.HealthCheckType,
		"instances":                  app.Instances,
		"memory":                     app.Memory,
		"name":                       app.Name,
		"package_state":              app.PackageState,
		"package_updated_at":         app.PackageUpdatedAt,
		"ports":                      app.Ports,
		"production":                 app.Production,
		"route_mappings_url":         fmt.Sprintf("/v2/apps/%s/route_mappings", app.Guid),
		"routes_url":                 fmt.Sprintf("/v2/apps/%s/routes", app.Guid),
		"service_bindings_url":       fmt.Sprintf("/v2/apps/%s/service_bindings", app.Guid),
		"space_guid":                 app.SpaceGuid,
		"space_url":                  fmt.Sprintf("/v2/spaces/%s", app.SpaceGuid),
		"stack_guid":                 "not-yet-implemented",
		"stack_url":                  "/v2/stacks/not-yet-implemented",
		"staging_failed_description": app.StagingFailedDescription,
		"staging_failed_reason":      app.StagingFailedReason,
		"staging_task_id":            app.StagingTaskID,
		"state":                      app.State,
		"version":                    "0",
	}
}
