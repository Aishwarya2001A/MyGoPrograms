package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wgIns sync.WaitGroup

func main() {

	counter := 0

	wgIns.Add(10)
	for i := 0; i < 10; i++ {

		go func() {
			for j := 0; j < 10; j++ {

				//time.Sleep(time.Millisecond * 1)

				counter += 1

			}
			wgIns.Done()
		}()
	}

	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())

	wgIns.Wait()

	fmt.Println("Counter value = ", counter)

	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())

}
