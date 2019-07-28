package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (a *API) PutRouteApp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//TODO: implement
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{})
}
