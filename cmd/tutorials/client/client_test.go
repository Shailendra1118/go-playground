package client

import (
	//"errors"
	"fmt"
	"go-playground/cmd/tutorials/users"
	"testing"

	"github.com/stretchr/testify/assert"
)
func TestGetUserWithId(t *testing.T) {
	fmt.Println("Testing...")

	mockUserGetter := &MockUserGetter{
		User : users.User{Name: "Aman", Addr: "London"},
		Err  :  nil, //errors.New("user not found"),
	}

	//client := &Client{}
	// Create a Client with the mock UserGetter
	client := Client{UserGetter: mockUserGetter}
	u, err := client.GetUserWithId("Shailendra")	
	assert.Nil(t, err)
	assert.Equal(t, u, users.User{Name: "Aman", Addr: "London"},)

}
