package main

import (
	"fmt"
)

func main() {
	greeting("Alex")
	fmt.Println(mtultiple(2, 3))
}

func greeting(name string) {
	fmt.Printf("Hello %v \n", name)
}

func mtultiple(a int, b int) int {
	return a * b
}
