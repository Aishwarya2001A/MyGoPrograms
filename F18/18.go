package main

import (
	"bytes"
	"fmt"
)

func main() {

	slice_1 := []byte{'A', 'i', 's', 'h', 'U'}
	slice_2 := []byte{'a', 'i', 'H', 'H', 'u'}

	res := bytes.Compare(slice_1, slice_2)
	if res == 0 {
		fmt.Println("Slices are equal")
	} else {
		fmt.Println("Slice are not equal")
	}
}
