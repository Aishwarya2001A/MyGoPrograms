package main

import (
	"fmt"
	"strings"
)

func main() {

	var message1 = "Hello,"

	message2 := "Welcome to Progress"

	fmt.Println(message1)
	fmt.Println(message2)

	message3 := `I love Go Programming`

	fmt.Println(message3)

	fmt.Printf("%c\n", message3[0])

	fmt.Printf("%c\n", message3[3])

	//stringLength := len(message2)

	fmt.Println("Length of a string is:", len(message2))

	result := message1 + " " + message2

	fmt.Println(result)

	string1 := "hi"
	string2 := "hello"
	string3 := "bye"

	fmt.Println(strings.Compare(string1, string2))
	fmt.Println(strings.Compare(string2, string3))

	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	result1 := strings.Contains(text, substring1)
	fmt.Println(result1)

	result2 := strings.Contains(text, substring2)
	fmt.Println(result2)
	text1 := "car"
	fmt.Println("Old String:", text1)

	replacedText := strings.Replace(text1, "r", "t", 1)

	fmt.Println("New String:", replacedText)

	text2 := "go is fun to learn"

	text2 = strings.ToUpper(text2)

	fmt.Println(text2)

	text3 := "I LOVE GOLANG"

	text3 = strings.ToLower(text3)
	fmt.Println(text3)

}
