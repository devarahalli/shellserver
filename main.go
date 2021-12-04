package main

import (
	"fmt"
	"shellserver/shellserver"
)

// Start a web server
// Define API to receive a command from the user
// Autenticate the user using the Basic authentication method.
// a. Supported authentication method is 1. Basic 2. Vault token based.
// Verify the user group.
// Check the command in the policiy server before allowing the command execution from the user.
//Check the user is

func main() {
	fmt.Println("Starting server on http://127.0.0.1:8080")
	shellserver.StartHttpServer(":8080")
}
