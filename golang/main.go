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

	names2 := names
	names3 := &names
	fmt.Println(names[0])
	fmt.Println(names2[0])
	fmt.Println(names3[0])
	names2[0] = "Changed"
	fmt.Println(names2[0])
	fmt.Println(names[0])
	names3[0] = "changed"
	fmt.Println(names3[0])
	fmt.Println(names[0])

}
