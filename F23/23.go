package main

import "fmt"

func main() {

	//the order of execution is LIFO. That is, the last defer statement is executed first,
	//and the first defer statement is executed last

	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")

	defer fmt.Println("hi")
	defer fmt.Println("hello")
	defer fmt.Println("bye")

}
