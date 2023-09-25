package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString("Hi this is Aishwarya")
	file.Close()
	stream, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
