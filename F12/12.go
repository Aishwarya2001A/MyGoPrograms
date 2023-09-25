package main

// import the errors package
import (
	"errors"
	"fmt"
)

func checkName(name string) error {

	newError := errors.New("Invalid Name")

	if name != "Aishwarya" {
		return newError
	}

	return nil
}

func main() {

	name := "Hello"

	err := checkName(name)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Valid Name")
	}

}
