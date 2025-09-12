package main

import (
	"fmt"
	"library/book"
)

func main() {
	fmt.Println("Hello, library!")
	mobyDick := book.CreateBook(
		"Momo",
		"Michel Ende",
		300,
	)
	mobyDick.PrintBook()
	fmt.Println(mobyDick.GetTitle())
	mobyDick.SetTitle("Moby Dick - special edition")
	mobyDick.PrintBook()

	algoritmoBook := book.NewTextBook(
		"Algoritmos",
		"Erika Flores",
		400,
		"Porrua",
		"advanced",
	)
	algoritmoBook.PrintTextBook()
}
