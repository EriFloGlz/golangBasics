package main

import "fmt"

func main() {
	hello()

	user := helloUser("totoro")
	fmt.Println(user)

	sum, mul := operations(3, 2)

	fmt.Println("La suma es :", sum)

	fmt.Println("La multiplicacion es :", mul)

}
func hello() {
	fmt.Println("Hello")
}
func helloUser(name string) string {
	return "Hello " + name
}
func operations(x, y int) (sum, mul int) {
	sum = x + y
	mul = x * y
	return
}
func getUser() (string, string) {

}
