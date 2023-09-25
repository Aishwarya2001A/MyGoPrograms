package main

import (
	"fmt"
	"reflect"
)

func main() {

	rune1 := 'A'
	rune2 := 'a'
	rune3 := '\a'

	fmt.Printf("Rune 1: %c; Unicode: %U; Type: %s", rune1,
		rune1, reflect.TypeOf(rune1))

	fmt.Printf("\nRune 2: %c; Unicode: %U; Type: %s", rune2,
		rune2, reflect.TypeOf(rune2))

	fmt.Printf("\nRune 3: Unicode: %U; Type: %s", rune3,
		reflect.TypeOf(rune3))

}
