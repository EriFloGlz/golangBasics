package main

import (
	"fmt"
)

func main() {
	//array
	pokemon := [...]string{"charmander", "pikachu", "raychu"}
	//fmt.Println(pokemon)
	for i, v := range pokemon {
		fmt.Printf("index : %d , value : %s\n", i, v)
	}
	x := len(pokemon)
	fmt.Print(x)

	//slice
	pokemonTrainer := []string{"Ash", "Alain", "Alder", "Jitsi"}
	a := pokemonTrainer[1:3]
	var pokemonEvolution [10]string
	fmt.Println(pokemonEvolution)
	fmt.Println(a)
	fmt.Println(pokemonTrainer)
	fmt.Println(len(pokemonTrainer))
	fmt.Println(cap(pokemonTrainer))
	all := append(pokemonTrainer, "square", "Dito")
	fmt.Println(all)
	score := make([]float64, 10)
	fmt.Print(score)
	for _, v := range score {
		fmt.Print(v)
	}
}
