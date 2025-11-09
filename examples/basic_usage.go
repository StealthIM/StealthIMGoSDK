package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/StealthIM/StealthIMGoSDK/stealthim"
)

func main() {
	// Get server URL from environment variable
	serverURL := os.Getenv("STEALTHIM_SERVER_URL")

	// Create a new server instance
	server := stealthim.NewServer(serverURL)

	// Ping the server to check if it's available
	ctx := context.Background()
	if err := server.Ping(ctx); err != nil {
		log.Fatalf("Failed to ping server: %v", err)
	}
	fmt.Println("Server is available!")

	// Register a new user (uncomment to use)
	/*
		if err := server.Register(ctx, "testuser", "password123", "Test User", "test@example.com", "1234567890"); err != nil {
			log.Fatalf("Failed to register user: %v", err)
		}
		fmt.Println("User registered successfully!")
	*/

	// Login with existing user (uncomment to use)
	/*
		user, userInfo, err := server.Login(ctx, "testuser", "password123")
		if err != nil {
			log.Fatalf("Failed to login: %v", err)
		}
		fmt.Printf("User logged in: %+v\n", userInfo)

		// Get user's own info
		selfInfo, err := user.GetSelfInfo(ctx)
		if err != nil {
			log.Fatalf("Failed to get self info: %v", err)
		}
		fmt.Printf("Self info: %+v\n", selfInfo)
	*/
}