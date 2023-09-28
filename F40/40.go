package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
func portal1(channel1 chan string) {
	for i := 0; i <= 3; i++ {
		channel1 <- "Welcome to channel 1"
	}

}

func portal2(channel2 chan string) {
	channel2 <- "Welcome to channel 2"
}

func Aname() {

	arr1 := [4]string{"Rohit", "Suman", "Aman", "Ria"}

	for t1 := 0; t1 <= 3; t1++ {

		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%s\n", arr1[t1])
	}
}

func Aid() {

	arr2 := [4]int{300, 301, 302, 303}

	for t2 := 0; t2 <= 3; t2++ {

		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d\n", arr2[t2])
	}
}

func main() {

	go display("Welcome")

	display("Hello")

	go func() {

		fmt.Println("Welcome!! to Progress")
	}()
	R1 := make(chan string)
	R2 := make(chan string)

	go portal1(R1)
	go portal2(R2)

	select {
	case op1 := <-R1:
		fmt.Println(op1)
	case op2 := <-R2:
		fmt.Println(op2)

	}
	go Aname()

	go Aid()

	time.Sleep(3500 * time.Millisecond)
}
