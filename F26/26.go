package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Printf("a: Type: %T, Length: %d, Capacity: %d\n", a, len(a), cap(a))
	fmt.Println("value of a:", a)

	b := make([]float64, 10, 20)
	fmt.Printf("b: Type: %T, Length: %d, Capacity: %d\n", b, len(b), cap(b))
	fmt.Println("value of b:", b)

	c := make([]string, 0, 5)
	fmt.Printf("c: Type: %T, Length: %d, Capacity: %d\n", c, len(c), cap(c))
	fmt.Println("value of c:", c)

	var student = make(map[string]int)

	student["Alvin"] = 21
	student["Alex"] = 47
	student["Mark"] = 27

	fmt.Println(student)
	fmt.Printf("Type: %T, Length: %d\n",
		student, len(student))

}
