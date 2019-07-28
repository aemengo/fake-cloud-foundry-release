package db

import (
	"fmt"
	"math/rand"
)

type User struct {
	Name     string
	Password string
	Token    string
}

func (db *DB) loadUsers() {
	newToken := func() string{
		b := make([]byte, 10)
		rand.Read(b)
		return fmt.Sprintf("%x", b)
	}

	for _, user := range db.config.Users {
		db.users = append(db.users, User{
			Name:     user.Name,
			Password: user.Password,
			Token:    newToken(),
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