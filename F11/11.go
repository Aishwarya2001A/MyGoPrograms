package main

import "fmt"

func main() {

	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}
}
