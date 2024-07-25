package main

import (
	"fmt"
)

func main() {

	simpleint := 10
	simplestring := "Hello"
	var intpointer *int = &simpleint
	var stringpointer *string = &simplestring
	fmt.Println(intpointer)
	fmt.Println(stringpointer)
	fmt.Println(*intpointer)
	fmt.Println(*stringpointer)

}
