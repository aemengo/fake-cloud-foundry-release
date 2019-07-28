package api

import (
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (a *API) Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/v2/info", a.Info)
	router.GET("/v2/organizations", Auth(a.database, a.Orgs))
	router.GET("/v2/organizations/:guid/spaces", Auth(a.database, a.OrgSpaces))
	return router
}

func Auth(db *db.DB, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Authorization")

		_, ok := db.GetUserByToken(token)
		if !ok {
			http.Error(w, `{"name":"UnauthorizedError","message":"You are not authorized to perform the requested action"}`, http.StatusUnauthorized)
			return
		}

		h(w, r, ps)
	}
}
