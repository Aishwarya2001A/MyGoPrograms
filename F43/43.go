// with race codition different outouts for each trails
package main

import (
	"fmt"
	"sync"
)

var aish = 0

func worker(wg *sync.WaitGroup) {
	aish = aish + 1

	wg.Done()
}
func main() {

	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go worker(&w)
	}

	w.Wait()
	fmt.Println("Value of x", aish)
}
