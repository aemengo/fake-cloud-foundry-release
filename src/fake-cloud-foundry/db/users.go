package db

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Name     string
	Password string
	Token    string
}

//{
//  "jti": "9e26dc08d72e4742b6cbcd993b5c4317",
//  "sub": "b824173d-b778-4d02-9625-84b1e5d1d28a",
//  "scope": [
//    "clients.read",
//    "openid",
//    "routing.router_groups.write",
//    "scim.read",
//    "cloud_controller.admin",
//    "uaa.user",
//    "routing.router_groups.read",
//    "cloud_controller.read",
//    "password.write",
//    "cloud_controller.write",
//    "network.admin",
//    "doppler.firehose",
//    "scim.write"
//  ],
//  "client_id": "cf",
//  "cid": "cf",
//  "azp": "cf",
//  "grant_type": "password",
//  "user_id": "b824173d-b778-4d02-9625-84b1e5d1d28a",
//  "origin": "uaa",
//  "user_name": "admin",
//  "email": "admin",
//  "auth_time": 1563827006,
//  "rev_sig": "8bfb1050",
//  "iat": 1563827006,
//  "exp": 1563827606,
//  "iss": "https://uaa.dev.cfdev.sh/oauth/token",
//  "zid": "uaa",
//  "aud": [
//    "scim",
//    "cloud_controller",
//    "password",
//    "cf",
//    "clients",
//    "uaa",
//    "openid",
//    "doppler",
//    "routing.router_groups",
//    "network"
//  ]
//}

func (db *DB) loadUsers() {
	newToken := func(name string) string{
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			 "jti": "jti",
			 "sub": name,
			 "scope": []string{"admin"},
			 "client_id": "cf",
			 "cid": "cf",
			 "azp": "cf",
			 "grant_type": "password",
			 "user_id": name,
			 "origin": "uaa",
			 "user_name": name,
			 "email": name,
			 "auth_time": time.Now().Unix(),
			 "rev_sig": "0",
			 "iat": time.Now().Unix(),
			 "exp": time.Now().Add(time.Hour * 24 * 24).Unix(),
			 "iss": fmt.Sprintf("https://%s/oauth/token", db.config.Domain()),
			 "zid": "uaa",
			 "aud": []string{"all"},
		})

		token, _ := jwtToken.SignedString([]byte("token-secret"))
		return token
	}

	for _, user := range db.config.Users {
		db.users = append(db.users, User{
			Name:     user.Name,
			Password: user.Password,
			Token:    newToken(user.Name),
		})
	}
}

func (db *DB) GetUserByNameAndPassword(username string, password string) (User, bool) {
	for _, user := range db.users {
		if user.Name == username && user.Password == password {
			return user, true
		}
	}

	return User{}, false
}

func (db *DB) GetUserByToken(token string) (User, bool) {
	for _, user := range db.users {
		if user.Token == token {
			return user, true
		}
	}

	return User{}, false
}