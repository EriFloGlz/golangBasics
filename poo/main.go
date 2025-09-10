package main

import (
	"fmt"
	"library/user"
)

func main() {
	newUser := user.User{
		Name: "Eri",
	} // Assuming User is a struct in the user package
	fmt.Println(newUser.Greet())
}
