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

	day_of_week := 6
	switch day_of_week {
	case 1:
		fmt.Println("Today is Saturday")
	case 2:
		fmt.Println("Today is Sunday")
	case 3:
		fmt.Println("Today is Monday")
	case 4:
		fmt.Println("Today is Tuesday")
	case 5:
		fmt.Println("Today is Wednesday")
	case 6:
		fmt.Println("Today is Friday")
	default:
		fmt.Println("Today number is invalid")
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		} else if i == 9 {
			break
		} else {
			fmt.Println(i)
		}
	}

}
