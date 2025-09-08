package greetings

import "fmt"

func Hello(name string) string {
	greeting := fmt.Sprintf("Hello %v ¡Welcome to the jungle!", name)
	return greeting
}
