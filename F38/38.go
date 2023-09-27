package main

import (
	"fmt"
	"reflect"
)

type mobile struct {
	price float64
	color string
}

type T struct {
	A int
	B string
	C float64
	D bool
}

//type Sum func(int64, int64) int64

func main() {
	destination := reflect.ValueOf([]string{"A", "B", "C"})
	source := reflect.ValueOf([]string{"D", "E", "F"})

	// Copy() function is used and it returns the number of elements copied
	counter := reflect.Copy(destination, source)
	fmt.Println(counter)

	fmt.Println(source)
	fmt.Println(destination)

	s1 := []string{"A", "B", "C", "D", "E"}
	s2 := []string{"D", "E", "F"}
	result := reflect.DeepEqual(s1, s2)
	fmt.Println(result)

	// DeepEqual is used to check two arrays are equal or not
	n1 := [5]int{1, 2, 3, 4, 5}
	n2 := [5]int{1, 2, 3, 4, 5}
	result = reflect.DeepEqual(n1, n2)
	fmt.Println(result)

	// DeepEqual is used to check two structures are equal or not
	m1 := mobile{500.50, "red"}
	m2 := mobile{400.50, "black"}
	result = reflect.DeepEqual(m1, m2)
	fmt.Println(result)

	theList := []int{1, 2, 3, 4, 5}
	swap := reflect.Swapper(theList)
	fmt.Printf("Original Slice :%v\n", theList)

	// Swapper() function is used to swaps the elements of slice
	swap(1, 3)
	fmt.Printf("After Swap :%v\n", theList)

	// Reversing a slice using Swapper() function
	for i := 0; i < len(theList)/2; i++ {
		swap(i, len(theList)-1-i)
	}
	fmt.Printf("After Reverse :%v\n", theList)
	v1 := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(v1))

	v2 := "Hello World"
	fmt.Println(reflect.ValueOf(v2))

	t := T{10, "ABCD", 15.20, true}
	typeT := reflect.TypeOf(t)
	fmt.Println(typeT.NumField())

	t = T{10, "ABCD", 15.20, true}
	typeT = reflect.TypeOf(t)

	for i := 0; i < typeT.NumField(); i++ {
		field := typeT.Field(i)
		fmt.Println(field.Name, field.Type)
	}

	var str []string
	var strType reflect.Value = reflect.ValueOf(&str)
	newSlice := reflect.MakeSlice(reflect.Indirect(strType).Type(), 10, 15)

	fmt.Println("Kind :", newSlice.Kind())
	fmt.Println("Length :", newSlice.Len())
	fmt.Println("Capacity :", newSlice.Cap())

	var str1 map[string]string
	var str1Type reflect.Value = reflect.ValueOf(&str1)
	newMap := reflect.MakeMap(reflect.Indirect(str1Type).Type())

	fmt.Println("Kind :", newMap.Kind())

}
