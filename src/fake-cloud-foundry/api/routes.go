package api

import "github.com/julienschmidt/httprouter"

func (a *API) Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/v2/info", a.Info)
	return router
}