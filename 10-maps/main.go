package main

import "fmt"

type square struct {
	x, y float64
}

func main() {
	var m = map[string]square{
		"red":  {40.32, 12.50},
		"blue": {15.3, 7.50},
	}
	store := map[string]float64{
		"apple":   12.65,
		"avocado": 7.21,
		"pinaple": 6.5,
	}
	store["grapes"] = 23.50
	fmt.Println(m["red"])
	fmt.Println(store["pinaple"])
	fmt.Println(store)

	_, ok := store["tomates"]

	if !ok {
		fmt.Print("El valor no existe")

	}
	for key, value := range store {
		fmt.Printf("key %s, - %f\n", key, value)
	}

}
