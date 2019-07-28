package uaa

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

//Sample response
//
//{
//  "app": {"version":"4.31.0-SNAPSHOT"},
//	"links":{
//		"uaa":"https://uaa.dev.cfdev.sh",
//		"passwd":"/forgot_password",
//		"login":"https://login.dev.cfdev.sh",
//		"register":"/create_account"
//	},
//	"zone_name":"uaa",
//	"entityID":"login.dev.cfdev.sh",
//	"commit_id":"2e27fa7",
//	"idpDefinitions":{},
//	"prompts":{
//		"username":["text","Email"],
//		"password":["password","Password"]},
//		"timestamp":"2019-03-15T19:54:53+0000"
//}


func (u *UAA) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	info := map[string]interface{}{
		"app": map[string]interface{}{
			"version": "0",
		},
		"links": map[string]interface{}{
			"uaa": fmt.Sprintf("http://uaa.%s", u.config.Domain()),
			"passwd": "/forgot_password",
			"login": fmt.Sprintf("http://uaa.%s", u.config.Domain()),
			"register": "/create_account",
		},
		"zone_name": "uaa",
		"entityID": fmt.Sprintf("uaa.%s", u.config.Domain()),
		"commit_id": "0",
		"idpDefinitions": map[string]interface{}{},
		"prompts": map[string]interface{}{
			"username": []string{"text", "Email"},
			"password": []string{"password","Password"},
		},
		"timestamp": time.Now().Format(time.RFC3339),
	}

	json.NewEncoder(w).Encode(info)
}


