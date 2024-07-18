package client

import (
	//"errors"
	"fmt"
	"go-playground/cmd/tutorials/users"
)

// MockUserGetter is a mock implementation of the UserGetter interface for testing
type MockUserGetter struct {
	User      users.User
	Err       error
}

// GetUser mocks the GetUser method
func (m *MockUserGetter) GetUser(id string) (users.User, error) {
	fmt.Println("Returning from GetUser of MockUserGetter")
	return m.User, m.Err
}