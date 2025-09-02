package main

import (
	"fmt"
	"math/rand/v2"
)

type User struct {
	ID   int
	Name string
}

func (u *User) GetPermission(status string) string {
	if status == "admin" {
		return "You have permission"
	}
	return "You don't have permission"
}

func main() {

	chucho := User{
		ID:   1,
		Name: "Chucho",
	}
	fmt.Println(chucho.ID)
	fmt.Println(chucho.Name)

	permission := chucho.GetPermission("user")
	fmt.Println(permission)
	numberInput := rand.IntN(100)
	fmt.Println(numberInput)

	for i := 0; i < 10; i++ {
		var numberUser int
		fmt.Print("Insert a number")
		fmt.Scanln(&numberUser)
		if numberInput == numberUser {
			fmt.Println("You guessed the number")
			break
		} else {
			fmt.Println("You didn't guess the number")
		}
	}
}
