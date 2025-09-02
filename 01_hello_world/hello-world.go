package main

import "fmt"

func main() {
	fmt.Println("hello world")

	//inicialiaion
	var x int
	var isPossible bool

	var a, b int
	a = 10
	b = 20
	fmt.Print(a + b)
	//declaration
	x = 10
	isPossible = true

	fmt.Println(x)
	fmt.Println(isPossible)

	// declaration and initialization
	var age, height int = 20, 200
	var userName, isMarried = "John", false
	fmt.Println(userName)
	fmt.Println(isMarried)
	fmt.Println(age)
	fmt.Println(height)

	//shorthand
	name := "John"
	fmt.Println(name)

}
