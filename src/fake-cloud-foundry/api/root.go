package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//{
//  "links": {
//    "app_ssh": {
//      "href": "ssh.dev.cfdev.sh:2222",
//      "meta": {
//        "host_key_fingerprint": "96:4d:89:2d:39:18:bc:16:e1:d3:d8:44:f8:16:af:85",
//        "oauth_client": "ssh-proxy"
//      }
//    },
//    "bits_service": null,
//    "cloud_controller_v2": {
//      "href": "https://api.dev.cfdev.sh/v2",
//      "meta": {
//        "version": "2.133.0"
//      }
//    },
//    "cloud_controller_v3": {
//      "href": "https://api.dev.cfdev.sh/v3",
//      "meta": {
//        "version": "3.68.0"
//      }
//    },
//    "credhub": null,
//    "logging": {
//      "href": "wss://doppler.dev.cfdev.sh:443"
//    },
//    "network_policy_v0": {
//      "href": "https://api.dev.cfdev.sh/networking/v0/external"
//    },
//    "network_policy_v1": {
//      "href": "https://api.dev.cfdev.sh/networking/v1/external"
//    },
//    "routing": {
//      "href": "https://api.dev.cfdev.sh/routing"
//    },
//    "self": {
//      "href": "https://api.dev.cfdev.sh"
//    },
//    "uaa": {
//      "href": "https://uaa.dev.cfdev.sh"
//    }
//  }
//}

func (a *API) Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := map[string]interface{}{
		"links": map[string]interface{}{
			"app_ssh": map[string]interface{}{
				"href": fmt.Sprintf("ssh.%s", a.config.Domain()),
				"meta": map[string]interface{}{
					"host_key_fingerprint": "96:4d:89:2d:39:18:bc:16:e1:d3:d8:44:f8:16:af:85",
					"oauth_client":         "ssh-proxy",
				},
			},
			"bits_service": nil,
			"cloud_controller_v2": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v2", a.config.Domain()),
				"meta": map[string]interface{}{
					"version": "999.999.999",
				},
			},
			"cloud_controller_v3": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3", a.config.Domain()),
				"meta": map[string]interface{}{
					"version": "999.999.999",
				},
			},
			"credhub": nil,
			"loggging": map[string]interface{}{
				"href": fmt.Sprintf("wss://doppler.%s", a.config.Domain()),
			},
			"network_policy_v0": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/networking/v0/external", a.config.Domain()),
			},
			"network_policy_v1": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/networking/v1/external", a.config.Domain()),
			},
			"routing": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/routing", a.config.Domain()),
			},
			"self": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s", a.config.Domain()),
			},
			"uaa": map[string]interface{}{
				"href": fmt.Sprintf("http://uaa.%s", a.config.Domain()),
			},
		},
	}

	json.NewEncoder(w).Encode(info)
}

//{
//  "links": {
//    "apps": {
//      "href": "https://api.dev.cfdev.sh/v3/apps"
//    },
//    "buildpacks": {
//      "experimental": true,
//      "href": "https://api.dev.cfdev.sh/v3/buildpacks"
//    },
//    "builds": {
//      "href": "https://api.dev.cfdev.sh/v3/builds"
//    },
//    "deployments": {
//      "experimental": true,
//      "href": "https://api.dev.cfdev.sh/v3/deployments"
//    },
//    "domains": {
//      "experimental": true,
//      "href": "https://api.dev.cfdev.sh/v3/domains"
//    },
//    "droplets": {
//      "href": "https://api.dev.cfdev.sh/v3/droplets"
//    },
//    "feature_flags": {
//      "href": "https://api.dev.cfdev.sh/v3/feature_flags"
//    },
//    "isolation_segments": {
//      "href": "https://api.dev.cfdev.sh/v3/isolation_segments"
//    },
//    "organizations": {
//      "href": "https://api.dev.cfdev.sh/v3/organizations"
//    },
//    "packages": {
//      "href": "https://api.dev.cfdev.sh/v3/packages"
//    },
//    "processes": {
//      "href": "https://api.dev.cfdev.sh/v3/processes"
//    },
//    "resource_match": {
//      "experimental": true,
//      "href": "https://api.dev.cfdev.sh/v3/resource_match"
//    },
//    "self": {
//      "href": "https://api.dev.cfdev.sh/v3"
//    },
//    "service_instances": {
//      "href": "https://api.dev.cfdev.sh/v3/service_instances"
//    },
//    "spaces": {
//      "href": "https://api.dev.cfdev.sh/v3/spaces"
//    },
//    "stacks": {
//      "href": "https://api.dev.cfdev.sh/v3/stacks"
//    },
//    "tasks": {
//      "href": "https://api.dev.cfdev.sh/v3/tasks"
//    }
//  }
//}

func (a *API) V3Root(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := map[string]interface{}{
		"links": map[string]interface{}{
			"apps": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/apps", a.config.Domain()),
			},
			"buildpacks": map[string]interface{}{
				"experimental": true,
				"href":         fmt.Sprintf("http://api.%s/v3/buildpacks", a.config.Domain()),
			},
			"builds": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/builds", a.config.Domain()),
			},
			"deployments": map[string]interface{}{
				"experimental": true,
				"href":         fmt.Sprintf("http://api.%s/v3/deployments", a.config.Domain()),
			},
			"domains": map[string]interface{}{
				"experimental": true,
				"href":         fmt.Sprintf("http://api.%s/v3/domains", a.config.Domain()),
			},
			"droplets": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/droplets", a.config.Domain()),
			},
			"feature_flags": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/feature_flags", a.config.Domain()),
			},
			"isolation_segments": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/isolation_segments", a.config.Domain()),
			},
			"organizations": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/organizations", a.config.Domain()),
			},
			"packages": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/packages", a.config.Domain()),
			},
			"processes": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/processes", a.config.Domain()),
			},
			"resource_match": map[string]interface{}{
				"experimental": true,
				"href":         fmt.Sprintf("http://api.%s/v3/resource_match", a.config.Domain()),
			},
			"self": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3", a.config.Domain()),
			},
			"service_instances": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/service_instances", a.config.Domain()),
			},
			"spaces": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/spaces", a.config.Domain()),
			},
			"stacks": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/stacks", a.config.Domain()),
			},
			"tasks": map[string]interface{}{
				"href": fmt.Sprintf("http://api.%s/v3/tasks", a.config.Domain()),
			},
		},
	}

	json.NewEncoder(w).Encode(info)
}
