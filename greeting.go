// Package main deofines the entry point of the Go application. It includes a function to greet the user by name.
package greeting

import "fmt"

// Greet takes a name as input and prints a greeting message to the console.
func Greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
