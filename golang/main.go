package main

import (
	"fmt"
)

func main() {
	greeting("Alex")
}

func greeting(name string) {
	fmt.Printf("Hello %v \n", name)
}
