package main

import (
	"fmt"
	"go-playground/cmd/tutorials/client"
	"go-playground/cmd/tutorials/users"
	//"slog"
)

func main() {
	//slog.Printf("Hello"), used with 1.21 onwards
	fmt.Printf("Starting...\n")
	//client.GetUserWithId("Shailendra Singh Yadav")

	userRepo := &users.UserRepo{}
	client := &client.Client{UserGetter: userRepo}

	//client is decouple from specific implementation of userRepo, I can have different UserGetter implmentation
	user, err := client.GetUserWithId("Aman")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	fmt.Printf("Found: %v\n", user)
	
}

