package main

import "fmt"

func main() {
	var username string = "Rıdvan"
	password := "pass123"
	const MaxLoginAttempts int = 3
	const pi float64 = 3.14

	fmt.Printf("Username: %s, Password: %s, Max Login Attempts: %d, Pi: %.2f\n", username, password, MaxLoginAttempts, pi)
}
