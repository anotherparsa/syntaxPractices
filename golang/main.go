package main

import "fmt"

func main() {

	var slice1 = []string{"E1", "E2", "E3", "E4"}
	fmt.Println(slice1[0])
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice2[3])
	var makeslice = make([]int, 4, 7)
	fmt.Println(makeslice)
	makeslice[0] = 0
	makeslice[1] = 1
	makeslice[2] = 2
	makeslice[3] = 3
	makeslice = append(makeslice, 4, 5, 6)
	fmt.Println(makeslice)

}
