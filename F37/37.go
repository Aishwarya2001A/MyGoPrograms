package main

import (
	"fmt"
	"strconv"
)

func main() {
	var strNum = "789"
	fmt.Printf("Before conversion:\n Type: %T\n Value: %v\n", strNum, strNum)
	val, err := strconv.Atoi(strNum)
	if err != nil {
		return
	}
	fmt.Printf("After conversion:\n Type: %T \n Value: %v\n", val, val)

	var num = 354
	fmt.Printf("Before conversion:\n Type: %T \n Value: %v\n", num, num)
	str := strconv.Itoa(num)
	fmt.Printf("After conversion:\n Type: %T\n Value: %v\n", str, str)

	bl := false
	fmt.Printf("Before conversion:\n Type: %T\n Value: %v\n", bl, bl)
	newBl := strconv.FormatBool(bl)
	fmt.Printf("After conversion:\n Type: %T\n Value: %v\n", newBl, newBl)

	blean := true
	fmt.Printf("Before conversion:\n Type: %T\n Value: %v\n", blean, blean)
	newBlean := strconv.FormatBool(blean)
	fmt.Printf("After conversion:\n Type: %T\n Value: %v\n", newBlean, newBlean)

}
