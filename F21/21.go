package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("unicode.In('x', unicode.Latin):", unicode.In('x', unicode.Latin))

	fmt.Println("unicode.In('F', unicode.ASCII_Hex_Digit):", unicode.In('F', unicode.ASCII_Hex_Digit))

	fmt.Println("unicode.In('\\t', unicode.White_Space):", unicode.In('\t', unicode.White_Space))

	fmt.Println("unicode.In('\\a', unicode.White_Space):", unicode.In('\a', unicode.White_Space))

	var r rune

	r = 'Q'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, r))
	fmt.Println()

	r = 'a'
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, r))
	fmt.Println()

	r = '='
	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, r))
	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, r))
	fmt.Println("bye")
}
