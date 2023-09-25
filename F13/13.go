package main

import "fmt"

func divide(num1, num2 int) error {

	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}

	return nil
}

func main() {

	err := divide(4, 0)

	if err != nil {
		fmt.Printf("error: %s", err)

	} else {
		fmt.Println("Valid Division")
	}
}
