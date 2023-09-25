package main

import "fmt"

func main() {

	var map_1 map[int]int

	if map_1 == nil {

		fmt.Println("True")
	} else {

		fmt.Println("False")
	}

	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}
	for id, pet := range map_2 {

		fmt.Println(id, pet)
	}

	fmt.Println("Map-2: ", map_2)
}
