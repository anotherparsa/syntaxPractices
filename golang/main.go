package main

import "fmt"

type address struct {
	state   string
	city    string
	zipcode int
}

func (a address) show() {
	fmt.Printf("Tthe state is %v and the city is %v and the zipcode is %v \n", a.state, a.city, a.zipcode)
}

func main() {
	var a address
	a.city = "Some City"
	a.state = "Some State"
	a.zipcode = 12345
	a.show()
}
