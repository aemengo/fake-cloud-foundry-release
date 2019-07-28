package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (a *API) PrivateDomains(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response := newResponse(nil)
	json.NewEncoder(w).Encode(response)
}
