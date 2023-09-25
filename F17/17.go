package main

import (
	"F16/Calculator"
	"fmt"
)

import "C"

func main() {
	fmt.Println(Calculator.Add(3, 5))
	fmt.Println(Calculator.Sub(5, 3))

}
