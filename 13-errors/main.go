package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("No se puedo dividir por 0")
	}
	return a / b, nil
}
func main() {
	result, err := divide(3, 0)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(result)
}
