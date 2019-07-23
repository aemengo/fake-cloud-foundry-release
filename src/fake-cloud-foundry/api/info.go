package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Sample response
//
//{
//	"name":"",
//    "build":"",
//    "support":"",
//    "version":0,
//    "description":"",
//    "authorization_endpoint":"https://login.dev.cfdev.sh",
//    "token_endpoint":"[PRIVATE DATA HIDDEN]",
//    "min_cli_version":null,
//    "min_recommended_cli_version":null,
//    "app_ssh_endpoint":"ssh.dev.cfdev.sh:2222",
//    "app_ssh_host_key_fingerprint":"96:4d:89:2d:39:18:bc:16:e1:d3:d8:44:f8:16:af:85",
//    "app_ssh_oauth_client":"ssh-proxy",
//    "doppler_logging_endpoint":"wss://doppler.dev.cfdev.sh:443",
//    "api_version":"2.133.0",
//    "osbapi_version":"2.14",
//    "routing_endpoint":"https://api.dev.cfdev.sh/routing"
//}

func (a *API) Info(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := map[string]interface{}{
		"name":                         "",
		"build":                        "",
		"support":                      "",
		"version":                      0,
		"description":                  "",
		//"authorization_endpoint":       fmt.Sprintf("http://login.%s", a.config.Domain), //<--- might remove login and just use UAA
		"authorization_endpoint":       "http://127.0.0.1:8081", //TODO fix
		"token_endpoint":               fmt.Sprintf("http://uaa.%s/oauth/token", a.config.Domain),
		"min_cli_version":              nil,
		"min_recommended_cli_version":  nil,
		"app_ssh_endpoint":             fmt.Sprintf("ssh.%s:2222", a.config.Domain),
		"app_ssh_host_key_fingerprint": "",
		"app_ssh_oauth_client":         "ssh-proxy",
		"doppler_logging_endpoint":     fmt.Sprintf("wss://doppler.%s:443", a.config.Domain),
		"api_version":                  "999.999.999",
		"osbapi_version":               "2.14",
		"routing_endpoint":             fmt.Sprintf("http://api.%s/routing", a.config.Domain), //<-- might change that to something different entirely
	}

	json.NewEncoder(w).Encode(info)
}
