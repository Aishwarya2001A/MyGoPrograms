package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string
	Age     int
	Address string
}

func main() {

	human1 := Human{"Aish", 22, "Mysore"}

	human_enc, err1 := json.Marshal(human1)

	if err1 != nil {

		fmt.Println(err1)
	}

	fmt.Println(string(human_enc))

	human2 := []Human{
		{Name: "Swati", Age: 23, Address: "New Delhi"},
		{Name: "Priya", Age: 20, Address: "Pune"},
		{Name: "Thanu", Age: 24, Address: "Bangalore"},
	}

	human2_enc, err2 := json.Marshal(human2)

	if err2 != nil {

		fmt.Println(err2)
	}

	fmt.Println()
	fmt.Println(string(human2_enc))

	var human3 Human

	Data := []byte(`{
        "Name": "Deeksha",  
        "Address": "Hyderabad",
        "Age": 21
    }`)

	err3 := json.Unmarshal(Data, &human3)

	if err3 != nil {

		fmt.Println(err3)
	}

	fmt.Println("Struct is:", human3)
	fmt.Printf("%s lives in %s.\n", human3.Name, human3.Address)

	var human4 []Human

	Data2 := []byte(`
    [
        {"Name": "Vani", "Address": "Delhi", "Age": 21},
        {"Name": "Rashi", "Address": "Noida", "Age": 24},
        {"Name": "Rohit", "Address": "Pune", "Age": 25}
    ]`)

	err3 = json.Unmarshal(Data2, &human4)

	if err2 != nil {

		fmt.Println(err3)
	}

	for i := range human4 {

		fmt.Println(human4[i])
	}
}
