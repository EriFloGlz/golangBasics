package main

import "fmt"

type Address struct {
	City, Country string
}
type Employee struct {
	Name    string
	Address Address
}

func (person *Employee) sayHi() {
	fmt.Printf("Hello I'm %s\n", person.Name)
}

func changeValue(x *int) {
	*x = 32
}
func main() {
	car := struct {
		Name string
		Year int
	}{
		Name: "Toyota",
		Year: 2025,
	}
	fmt.Println(car.Name)
	employeA := Employee{
		Name: "Erika",
		Address: Address{
			City:    "Monterrey",
			Country: "Mexico",
		},
	}
	employeA.sayHi()
	fmt.Println(employeA.Address.City)
	fmt.Println(employeA.Name)

	x := 10
	//pointerEmpoloye := &x

	//fmt.Println(pointerEmpoloye)
	//*pointerEmpoloye = 32
	//fmt.Println(*pointerEmpoloye)
	//fmt.Println(x)

	changeValue(&x)
	fmt.Println(x)

}
