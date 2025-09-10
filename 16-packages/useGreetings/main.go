package main

import (
	"fmt"
	"log"

	"github.com/daedraMex/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// show / hide hora y fecha
	//log.SetFlags(0)

	message, e := greetings.Hello("Anita")

	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(message)

	greeting := greetings.RandomNames("Montse")
	fmt.Println(greeting)

	names := []string{"lupita", "carmen"}
	sayHi, error := greetings.GreetingEveryone(names)
	if error != nil {
		log.Fatal(error)
	}
	for _, v := range sayHi {
		fmt.Println(v)
	}
}
