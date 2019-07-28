package api

import (
	"encoding/json"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (a *API) Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", a.Root)
	router.GET("/v2/info", a.Info)
	router.GET("/v2/organizations", Auth(a.database, a.Orgs))
	router.GET("/v2/organizations/:guid/spaces", Auth(a.database, a.OrgSpaces))
	router.GET("/v2/organizations/:guid/private_domains", Auth(a.database, a.PrivateDomains))
	router.GET("/v2/apps", Auth(a.database, a.GetApps))
	router.POST("/v2/apps", Auth(a.database, a.PostApps))
	router.GET("/v2/apps/:guid", Auth(a.database, a.GetApp))
	router.PUT("/v2/apps/:guid", Auth(a.database, a.PutApp))
	router.GET("/v2/apps/:guid/routes", Auth(a.database, a.AppRoutes))
	router.PUT("/v2/apps/:guid/bits", Auth(a.database, a.PutBits))
	router.GET("/v2/apps/:guid/stats", Auth(a.database, a.AppStats))
	router.GET("/v2/apps/:guid/instances", Auth(a.database, a.AppInstances))
	router.GET("/v2/routes", Auth(a.database, a.GetRoutes))
	router.POST("/v2/routes", Auth(a.database, a.PostRoutes))
	router.PUT("/v2/routes/:route_guid/apps/:app_guid", Auth(a.database, a.PutRouteApp))
	router.GET("/v2/jobs/:guid", Auth(a.database, a.Jobs))
	router.GET("/v2/shared_domains", Auth(a.database, a.SharedDomains))
	router.GET("/v2/shared_domains/:guid", Auth(a.database, a.SharedDomain))

	router.GET("/v3", a.V3Root)
	router.GET("/v3/apps", a.V3Apps)
	router.GET("/v3/apps/:guid/processes", a.V3AppProcesses)
	router.GET("/v3/apps/:guid/processes/:process", a.V3AppProcess)
	router.GET("/v3/apps/:guid/droplets/current", a.V3AppDroplets)
	router.GET("/v3/processes/:guid/stats", a.V3ProcessStats)

	router.NotFound = &notFoundHandler{}
	return router
}

func Auth(db *db.DB, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")

		_, ok := db.GetUserByToken(strings.Replace(token, "bearer ", "", 1))
		if !ok {
			http.Error(w, `{"name":"UnauthorizedError","message":"You are not authorized to perform the requested action"}`, http.StatusUnauthorized)
			return
		}

		h(w, r, ps)
	}
}

type notFoundHandler struct {}

func (h *notFoundHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code": 10000,
		"description": "Unknown request",
		"error_code": "CF-NotFound",
	})
}