package uaa

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
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

func (u *UAA) Token(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := map[string]interface{}{
		"access_token": "access-token",
		"token_type": "",
		"id_token": "id-token",
		"refresh_token": "refresh-token",
		"expires_in": 999999999,
		"scope":"clients.read openid routing.router_groups.write scim.read cloud_controller.admin uaa.user routing.router_groups.read cloud_controller.read password.write cloud_controller.write network.admin doppler.firehose scim.write",
		"jti": "jti",
	}

	json.NewEncoder(w).Encode(info)
}


