package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	// Create the tern command
	cmd := exec.Command(
		"tern", 
		"migrate", 
		"--migrations", 
		"./internal/store/pgstore/migrations", 
		"--config", 
		"./internal/store/pgstore/migrations/tern.conf")

	// Run the command
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Error running tern command: %v", err))
	}

	fmt.Println("Migration completed successfully.")
}
