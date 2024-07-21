package main

import (
	"fmt"
)

func main() {

	emptyMap := map[int]string{}
	emptyMap[1] = "One"
	emptyMap[2] = "Two"
	fmt.Println(emptyMap[1])

	initializedmap := map[int]string{1: "One", 2: "Two", 3: "Three", 4: "Four"}

	fmt.Println(initializedmap)
	//updating
	initializedmap[4] = "Chahar"
	fmt.Println(initializedmap)
	//adding
	initializedmap[5] = "Five"
	initializedmap[6] = "Six"
	fmt.Println(initializedmap)
	delete(initializedmap, 6)
	fmt.Println(initializedmap)

	for key, _ := range initializedmap {
		fmt.Println(initializedmap[key])
	}

}
