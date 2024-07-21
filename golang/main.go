package main

import (
	"fmt"
)

func main() {

	var names [3]string
	names[0] = "Name 1"
	names[1] = "Name 2"
	names[2] = "Name 3"

	//	numbers := [4]int{1, 2, 3, 4}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("The index is %v and the value is %v \n", index, value)
	}

}
