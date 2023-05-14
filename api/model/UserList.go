package model

import "fmt"

type UserList struct {
	list []User
}

func (ul UserList) Get() []User {
	return ul.list
}

func (ul *UserList) Append(user User) {
	ul.list = append(ul.list, user)
}

func (ul UserList) Find(match func(u User) bool) (User, error) {
	for _, user := range ul.list {
		if match(user) {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("can't find user")
}
