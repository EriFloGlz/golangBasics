package main

import (
	"fmt"
	"github.com/daedraMex/greetings"
)

func main() {
	message := greetings.Hello("Daedra")
	fmt.Println(message)
}
