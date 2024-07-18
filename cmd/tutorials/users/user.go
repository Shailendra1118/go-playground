package users

import (
	"fmt"
)

type User struct {
	Name string
	Addr string
}

// This interface is not needed
/*
   Go interfaces generally belong in the package that uses values of the interface type, 
   not the package that implements those values.
*/
// type UserRepository interface {
// 	GetAllUsers()
// 	GetUser(id string)
// }

type UserRepo struct {
	
}

func (UserRepo) GetAllUsers() {
	fmt.Printf("Getting all users...")
}

func (UserRepo) GetUser(id string) (User, error){
	fmt.Printf("Getting user with id: %s...\n", id)
	return User{Name: id, Addr: ""}, nil
}



