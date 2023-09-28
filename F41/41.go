package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "Aishu"
	}
	close(mychnl)
}
func sending(s chan<- string) {
	s <- "Aishwarya"
}

func main() {

	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)

	ch := make(chan int)

	go myfunc(ch)
	ch <- 23

	c := make(chan string)

	go myfun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)

		mychnl := make(chan string, 5)

		go func() {
			mychnl <- "hii"
			mychnl <- "helloo"
			mychnl <- "how are you"
			mychnl <- "bye"
			close(mychnl)
		}()

		for res := range mychnl {
			fmt.Println(res)
		}
		fmt.Println("Length of the channel is: ", len(mychnl))
		fmt.Println("Capacity of the channel is: ", cap(mychnl))

		mychanl1 := make(<-chan string)

		mychanl2 := make(chan<- string)

		fmt.Printf("%T", mychanl1)
		fmt.Printf("\n%T", mychanl2)

		chanl1 := make(chan string)

		go sending(chanl1)

		fmt.Println(<-chanl1)

	}
}
