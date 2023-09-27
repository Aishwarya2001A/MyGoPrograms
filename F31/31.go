package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {

	str := "Aishwarya"

	match1, err := regexp.MatchString("Aish", str)
	fmt.Println("Match: ", match1, " Error: ", err)

	str2 := "ComputerScience"
	match2, err := regexp.MatchString("com", str2)
	fmt.Println("Match: ", match2, "Error: ", err)

	match3, err := regexp.MatchString("ai(s", str2)
	fmt.Println("Match: ", match3, "Error: ", err)

	re, _ := regexp.Compile("aishu")

	str = "Hi i am aishu"

	match := re.FindStringIndex(str)
	fmt.Println(match)

	str2 = "I love golang"

	match = re.FindStringIndex(str2)
	fmt.Println(match)

	re2, _ := regexp.Compile("[0-9]+-a.*g")

	match = re2.FindStringSubmatchIndex("2001-aish_gowda")
	fmt.Println(match3)

	match4 := re.FindAllStringSubmatchIndex("I'am software engineer at"+
		"Progress", -1)
	fmt.Println(match4)

	re3, _ := regexp.Compile(" ")
	match5 := re3.ReplaceAllString("All I do"+
		" is code everytime.", "+")
	fmt.Println(match5)

	re4, _ := regexp.Compile("[aeiou]+")
	match6 := re4.ReplaceAllStringFunc("All I do"+
		" is code everytime.", strings.ToLower)
	fmt.Println(match6)
}
