package client

import (
	"fmt"
	"go-playground/cmd/tutorials/users"
)

type UserGetter interface {
	GetUser(id string) (users.User, error)
}

type Client struct {
	UserGetter UserGetter
}

func (c Client) GetUserWithId(id string) (users.User, error) {
	fmt.Printf("Fetch details of user with ID: %s\n", id)
	//userRepo := &users.UserRepo{}
	//user, err := userRepo.GetUser(id)
	user, err := c.UserGetter.GetUser(id)
	return user, err
}