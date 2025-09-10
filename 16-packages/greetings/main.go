package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("try with another string more longer")
	}
	greeting := fmt.Sprintf("Hello %v ¡Welcome to the jungle!", name)
	return greeting, nil
}
func RandomNames(name string) string {
	randomNames := []string{
		"Hello %v ¡Welcome to the jungle!",
		"Nice to see you %v",
		"Placer en conocerte %v",
		"Boun jour %v",
	}
	randomNumber := rand.Intn(len(randomNames))
	greeting := fmt.Sprintf(randomNames[randomNumber], name)
	return greeting
}
func GreetingEveryone(names []string) ([]string, error) {
	if len(names) == 0 {
		return nil, errors.New("names > 0 ")
	}
	newGreetings := []string{}
	for _, v := range names {
		newName := RandomNames(v)
		newGreetings = append(newGreetings, newName)

	}
	return newGreetings, nil
}
