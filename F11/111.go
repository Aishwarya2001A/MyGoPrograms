package main

// import the custom package calculator
import (
	"Calculator"
	"fmt"
)

func main() {

	number1 := 9
	number2 := 5

	fmt.Println(Calculator.Add(number1, number2))

	fmt.Println(Calculator.Subtract(number1, number2))

}
