package main

import (
	"fmt"
)

func main() {
	sum, avg := multiplereturn(2, 10)
	fmt.Printf("This is sum %v and this is average %v \n", sum, avg)
}

func multiplereturn(a int, b int) (int, int) {
	sum := a + b
	avg := sum / 2

	return sum, avg
}
