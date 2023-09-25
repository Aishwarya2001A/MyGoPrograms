package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"Aishwarya", "Mysore", 570017}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Swati", city: "Honnavar",
		Pincode: 570018}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Deepi"}
	fmt.Println("Address3: ", a3)

	fmt.Println("Name: ", a1.Name)
	fmt.Println("Name: ", a2.Name)
	fmt.Println("Name: ", a3.Name)

}
