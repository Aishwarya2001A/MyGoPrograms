package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := [7]string{"This", "is", "the", "Program",
		"of", "Go", "language"}

	fmt.Println("Array:", arr)

	myslice := arr[1:6]

	fmt.Println("Slice:", myslice)

	fmt.Printf("Length of the slice: %d", len(myslice))

	fmt.Printf("\nCapacity of the slice: %d", cap(myslice))

	var my_slice_1 = []string{"hi", "hello", "bye"}

	fmt.Println("My Slice 1:", my_slice_1)

	my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
	fmt.Println("My Slice 2:", my_slice_2)

	var my_slice_3 = make([]int, 4, 7)
	fmt.Printf("Slice 3 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_3, len(my_slice_3), cap(my_slice_3))

	var my_slice_4 = make([]int, 7)
	fmt.Printf("Slice 4= %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_4, len(my_slice_4), cap(my_slice_4))

	for e := 0; e < len(myslice); e++ {
		fmt.Println(myslice[e])
	}

	intSlice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

}
