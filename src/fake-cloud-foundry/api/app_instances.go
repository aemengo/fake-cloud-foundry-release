package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//{
//  "0": {
//    "since": 1563828463,
//    "state": "RUNNING",
//    "uptime": 0
//  }
//}

func (a *API) AppInstances(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	info := map[string]interface{}{
		"0": map[string]interface{}{
			"since":  0,
			"state":  "RUNNING",
			"uptime": 0,
		},
	}

	json.NewEncoder(w).Encode(info)
}
