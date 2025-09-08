package main

import (
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		panic("Woul be no divition by 0")
	}
	return a / b, nil

}
func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from", r)
			result = 0
		}
	}()
	if b == 0 {
		panic("Divide by 0 ")
	}
	return a / b
}
func main() {
	divide(2, 9)
	//divide(2, 0)
	safeDivide(4, 9)
	safeDivide(3, 0)
	safeDivide(9, 12)
}
