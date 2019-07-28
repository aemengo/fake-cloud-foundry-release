package db

import (
	"github.com/satori/go.uuid"
	"time"
)

type App struct {
	Guid                     string
	Name                     string
	SpaceGuid                string
	Buildpack                *string
	Command                  *string
	Console                  bool
	Debug                    *string
	DetectedBuildpack        *string
	DetectedBuildpackGuid    *string
	DetectedStartCommand     string
	DiskQuota                int
	EnableSSH                bool
	EnvironmentJSON          map[string]interface{}
	HealthCheckHTTPEndpoint  *string
	HealthCheckTimeout       *string
	HealthCheckType          string
	Instances                int
	Memory                   int
	PackageState             string
	PackageUpdatedAt         *string
	Ports                    []int
	Production               bool
	StagingFailedDescription *string
	StagingFailedReason      *string
	StagingTaskID            *string
	State                    string
	Version                  string
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

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

func (db *DB) GetApps() []App {
	return db.apps
}

func (db *DB) GetAppByGuid(guid string) (App, bool) {
	for _, app := range db.apps {
		if app.Guid == guid {
			return app, true
		}
	}

	return App{}, false
}

func (db *DB) GetAppByNameAndSpaceGuid(name string, spaceGuid string) (App, bool) {
	for _, app := range db.apps {
		if app.Name == name && app.SpaceGuid == spaceGuid {
			return app, true
		}
	}

	return App{}, false
}

func (db *DB) SaveApp(app App) {
	var (
		i = -1
	)

	for index, j := range db.apps {
		if j.Guid == app.Guid {
			i = index
		}
	}

	if i > -1 {
		app.UpdatedAt = time.Now().UTC()
		db.apps[i] = app
	}
}

func (db *DB) CreateApp(app App) App {
	now := time.Now().UTC()

	app.Guid = uuid.NewV4().String()
	app.Console = false
	app.DiskQuota = 1024
	app.EnableSSH = true
	app.DetectedStartCommand = ""
	app.EnvironmentJSON = map[string]interface{}{}
	app.HealthCheckType = "port"
	app.Instances = 1
	app.Memory = 256
	app.PackageState = "PENDING"
	app.Ports = []int{8080}
	app.Production = false
	app.State = "STOPPED"
	app.Version = "0"
	app.CreatedAt = now
	app.UpdatedAt = now

	db.apps = append(db.apps, app)
	return app
}
