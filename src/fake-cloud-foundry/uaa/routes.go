package uaa

import "github.com/julienschmidt/httprouter"

func (u *UAA) Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/login", u.Login)
	router.POST("/oauth/token", u.Token)
	return router
}
