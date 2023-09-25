package main

import "fmt"

func main() {

	var x int = 5748

	var p *int

	p = &x

	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)

	var s *int

	fmt.Println("s = ", s)

	var y = 458

	var p1 = &y

	fmt.Println("Value stored in y = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p1 = ", p1)

	z := 458

	q := &y

	fmt.Println("Value stored in z = ", z)
	fmt.Println("Address of z = ", &z)
	fmt.Println("Value stored in pointer variable q = ", q)

	var a = 458

	var b = &a

	fmt.Println("Value stored in y before changing = ", a)
	fmt.Println("Address of a = ", &a)
	fmt.Println("Value stored in pointer variable b= ", b)

	fmt.Println("Value stored in a(*b) Before Changing = ", *b)

	*b = 500

	fmt.Println("Value stored in a(*b) after Changing = ", b)
}
