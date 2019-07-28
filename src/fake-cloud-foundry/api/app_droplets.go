package api

import (
	"encoding/json"
	"fmt"
	cfg "github.com/aemengo/fake-cloud-foundry/config"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//{
//  "buildpacks": [
//    {
//      "buildpack_name": "ruby",
//      "detect_output": "ruby",
//      "name": "ruby_buildpack",
//      "version": "1.7.34"
//    }
//  ],
//  "checksum": {
//    "type": "sha256",
//    "value": "091e0275c457feb5ec703ece380e1e2155f38e06e72a1f75ce4ae0d5eb46b08f"
//  },
//  "created_at": "2019-07-22T20:47:26Z",
//  "error": null,
//  "execution_metadata": "",
//  "guid": "ba5206eb-9419-475b-8319-975bffa0b806",
//  "image": null,
//  "lifecycle": {
//    "data": {},
//    "type": "buildpack"
//  },
//  "links": {
//    "app": {
//      "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244"
//    },
//    "assign_current_droplet": {
//      "href": "https://api.dev.cfdev.sh/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/relationships/current_droplet",
//      "method": "PATCH"
//    },
//    "package": {
//      "href": "https://api.dev.cfdev.sh/v3/packages/e96e1d17-1be1-4dc5-a61d-b13bbdf31ce1"
//    },
//    "self": {
//      "href": "https://api.dev.cfdev.sh/v3/droplets/ba5206eb-9419-475b-8319-975bffa0b806"
//    }
//  },
//  "metadata": {
//    "annotations": {},
//    "labels": {}
//  },
//  "process_types": {
//    "web": "bundle exec rackup config.ru -p $PORT",
//    "worker": "bundle exec rackup config.ru"
//  },
//  "stack": "cflinuxfs3",
//  "state": "STAGED",
//  "updated_at": "2019-07-22T20:47:35Z"
//}

func (a *API) V3AppDroplets(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	response := presentAppDroplet(a.config)

	json.NewEncoder(w).Encode(response)
}

func presentAppDroplet(config cfg.Config) map[string]interface{} {
	return map[string]interface{}{
		"buildpacks": []map[string]interface{}{{
			"buildpack_name": "ruby",
			"detect_output":  "ruby",
			"name":           "ruby_buildpack",
			"version":        "1.7.34",
		}},
		"checksum": map[string]interface{}{
			"type":  "sha256",
			"value": "091e0275c457feb5ec703ece380e1e2155f38e06e72a1f75ce4ae0d5eb46b08f",
		},
		"created_at":         "2019-07-22T20:47:26Z",
		"error":              nil,
		"execution_metadata": "",
		"guid":               "ba5206eb-9419-475b-8319-975bffa0b806",
		"image":              nil,
		"lifecycle": map[string]interface{}{
			"data": map[string]interface{}{},
			"type": "buildpack",
		},
		"links": map[string]interface{}{
			"app": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244", config.Host),
			},
			"assign_current_droplet": map[string]interface{}{
				"href":   fmt.Sprintf("http://%s/v3/apps/cb05ab74-3a08-453b-ab14-62d0274c3244/relationships/current_droplet", config.Host),
				"method": http.MethodPatch,
			},
			"package": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/packages/e96e1d17-1be1-4dc5-a61d-b13bbdf31ce1", config.Host),
			},
			"self": map[string]interface{}{
				"href": fmt.Sprintf("http://%s/v3/droplets/ba5206eb-9419-475b-8319-975bffa0b806", config.Host),
			},
			"metadata": map[string]interface{}{
				"annotations": map[string]interface{}{},
				"labels":      map[string]interface{}{},
			},
			"process_types": map[string]interface{}{
				"web":    "bundle exec rackup config.ru -p $PORT",
				"worker": "bundle exec rackup config.ru",
			},
			"stack":      "cflinuxfs3",
			"state":      "STAGED",
			"updated_at": "2019-07-22T20:47:35Z",
		},
	}
}

