package main

import "fmt"

type tank interface {
	Tarea() float64
	Volume() float64
}

type myvalue struct {
	radius float64
	height float64
}

func (m myvalue) Tarea() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m myvalue) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func myfun(a interface{}) {

	switch b := a.(type) {

	case int:
		fmt.Println("Type: int, Value:", b)
	case string:
		fmt.Println("\nType: string, Value: ", b)
	case float64:
		fmt.Println("\nType: float64, Value: ", b)
	default:
		fmt.Println("\nType not found")
	}
}

func main() {

	var t tank
	t = myvalue{10.6, 14.99}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())
	myfun("Goprogramming")
	myfun(67.9)
	myfun(true)
}
