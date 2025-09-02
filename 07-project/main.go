package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	playGame()

}
func playGame() {
	intentos := 0
	maxIntentos := 4

	randomNumber := rand.Intn(10)
	fmt.Println("number", randomNumber)
	for intentos <= maxIntentos {
		intentos++
		var number int
		fmt.Println("Insert a number: ")
		fmt.Scanln(&number)
		if number == randomNumber {
			fmt.Println("Guess the number")
			playAgain()
		} else if number > randomNumber {
			fmt.Println("The number is greater than the random number")
		} else if number < randomNumber {
			fmt.Println("The number is less than the random number")
		}
	}
}
func playAgain() {
	fmt.Println("Do yo want to play again ?  (y/n)")
	var answer string
	fmt.Scanln(&answer)
	answerUppercase := strings.ToUpper(answer)
	switch answerUppercase {
	case "Y":
		playGame()
	case "N":
		fmt.Println("Thank you for playing")
	default:
		fmt.Println("Invalid answer")
		playAgain()
	}

}
