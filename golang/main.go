package main

import (
	"fmt"
)

func main() {

	age := 3
	is_healthy := false

	if age >= 18 {
		if is_healthy {
			fmt.Println("You're in legal age and healthy")
		} else {
			fmt.Println("You're in legal age but not healthy")
		}
	} else {
		if is_healthy {
			fmt.Println("Youre not in legal age but are healthy")
		} else {
			fmt.Println("You're neither in legal age nor healty")
		}
	}

}
