package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	gender string
}

func main() {

	person1 := new(Person)
	person2 := new(Person)

	person1.name = "Aish"
	person1.age = 22
	person1.gender = "female"

	person2.name = "Swati"
	person2.age = 21
	person2.gender = "female"

	fmt.Println(person1)
	fmt.Println(person2)

	type Slice []int

	slice1 := new(Slice)
	slice2 := new(Slice)

	tempSlice1 := append(*slice1, 1, 2, 3)
	slice1 = &tempSlice1

	tempSlice2 := append(*slice1, 7, 5, 3)
	slice2 = &tempSlice2

	fmt.Println(slice1)
	fmt.Println(slice2)

}
