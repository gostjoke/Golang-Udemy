package main

import (
	"fmt"
	"log"
)

type User struct {
	ID    int
	First string
}

type MockDatabse struct {
	Users map[int]User
}

func (md MockDatabse) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("user with id %v not found", id)
	}
	return user, nil
}

func (md MockDatabse) SaveUser(u User) error {
	_, ok := md.Users[u.ID]
	if ok {
		return fmt.Errorf("user with id %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

type DataStore interface {
	GetUser(id int) (User, error)
	SaveUser(u User) error
}

type Service struct {
	ds DataStore
}

func (s Service) GetUser(id int) (User, error) {
	return s.ds.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.ds.SaveUser(u)
}

func main() {
	db := MockDatabse{
		Users: make(map[int]User),
	}

	service := Service{ds: db}

	user := User{ID: 1, First: "John"}

	err := service.SaveUser(user)
	if err != nil {
		log.Fatalf("error saving user: %v", err)
	}

	u1Returned, err := service.GetUser(1)
	if err != nil {
		log.Fatalf("error getting user: %v", err)

	}
	fmt.Printf("User build: %+v\n", user)
	fmt.Printf("User retrieved: %+v\n", u1Returned)
}
