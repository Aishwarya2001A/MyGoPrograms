package main

import (
	"fmt"
	"os"
)

func main() {

	myProgramName := os.Args[0]

	cmdArgs := os.Args[4]

	gettingArgs := os.Args[2]

	toGetAllArgs := os.Args[1:]

	fmt.Println(myProgramName)

	fmt.Println(cmdArgs)

	fmt.Println(gettingArgs)

	fmt.Println(toGetAllArgs)
}
