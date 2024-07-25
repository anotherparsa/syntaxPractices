package main

import (
	"fmt"
)

func main() {
	a := 0
	b := 0
	fmt.Println(a)
	fmt.Println(b)
	passedByValue(a)
	passedByReference(&b)
	fmt.Println(a)
	fmt.Println(b)

}

func passedByValue(a int) {
	a = 10
}

func passedByReference(b *int) {
	*b = 10
}
