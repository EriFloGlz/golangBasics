package main

import "fmt"

const PI = 3.1416

func main() {
	var x int = 0

	for x <= 9 {
		fmt.Println(x)
		x++
	}
	for i := 1; i <= 6; i++ {
		if i == 2 {
			fmt.Println("salida")
			continue
		}
		fmt.Println(i)
	}
	/*
		Haciendo que itere en un rango de datos
	*/
	/*
		for i := range 6 {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}*/
	//Bucle infinito
	// for {
	// 	fmt.Println("Hola")
	// }

}
