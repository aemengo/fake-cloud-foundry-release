package uaa

import (
	"encoding/json"
	"fmt"
	"github.com/aemengo/fake-cloud-foundry/db"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

// Sample response
// {
//   "access_token":"[PRIVATE DATA HIDDEN]",
//   "token_type":"[PRIVATE DATA HIDDEN]",
//   "id_token":"[PRIVATE DATA HIDDEN]",
//   "refresh_token":"[PRIVATE DATA HIDDEN]",
//   "expires_in":599,
//   "scope":"clients.read openid routing.router_groups.write scim.read cloud_controller.admin uaa.user routing.router_groups.read cloud_controller.read password.write cloud_controller.write network.admin doppler.firehose scim.write",
//   "jti":"200b44ad1505413b96d56baeafde8903"
// }

func (u *UAA) Token(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()

	user, ok := u.user(r)
	if !ok {
		http.Error(w, `{"error":"unauthorized","error_description":"Bad credentials"}`, http.StatusUnauthorized)
		return
	}

	info := map[string]interface{}{
		"access_token":  user.Token,
		"id_token":      user.Token,
		"refresh_token": user.Token,
		"expires_in":    time.Now().Add(time.Hour * 24 * 24).Unix(),
		"token_type":    "bearer",
		"scope":         "admin",
		"jti":           "jti",
	}

	json.NewEncoder(w).Encode(info)
}

func (u *UAA) user(r *http.Request) (db.User, bool) {
	switch r.PostFormValue("grant_type") {
	case "password":
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		return u.database.GetUserByNameAndPassword(username, password)
	case "refresh_token":
		token := r.PostFormValue("refresh_token")
		return u.database.GetUserByToken(token)
	default:
		panic(fmt.Sprintf("unhandled token grant_type: %s", r.PostFormValue("grant_type")))
	}
}
