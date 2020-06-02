package models

import (
	"errors"
	"fmt"
)

type Users struct {
	Id        int
	FirstName string
	LastName  string
}

var (
	users  []*Users
	nextID = 1
)

func GetUser() []*Users {
	return users
}

func AddUser(u Users) (Users, error) {
	if u.Id != 0 {
		return Users{}, errors.New("new user must not include id")
	}
	u.Id = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (Users, error) {
	for _, u := range users {
		if u.Id == id {
			return *u, nil
		}
	}
	return Users{}, fmt.Errorf("user with id '%v' not found", id)
}

func UpdateUser(u Users) (Users, error) {
	for i, candidate := range users {
		if candidate.Id == u.Id {
			users[i] = &u
			return u, nil
		}
	}
	return Users{}, fmt.Errorf("user with id '%v' not found", u.Id)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.Id == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with id '%v' not found",id)
}
