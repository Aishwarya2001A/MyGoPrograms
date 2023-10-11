package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {

	aishu := &Person{
		Name: "Aishwarya",
		Age:  22,
	}

	data, err := proto.Marshal(aishu)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	newAishwarya := &Person{}
	err = proto.Unmarshal(data, newAishwarya)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newAishwarya.GetAge())
	fmt.Println(newAishwarya.GetName())
}
