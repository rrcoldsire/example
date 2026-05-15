// Package main deofines the entry point of the Go application. It includes a function to greet the user by name.
package main

import "fmt"

// Greeting takes a name as input and prints a greeting message to the console.
func Greeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
