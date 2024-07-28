package main

type address struct {
	state   string
	city    string
	zipcode int
}

func main() {
	var a address
	a.city = "Some City"
	a.state = "Some State"
	a.zipcode = 12345

	fmt.Println(a.zipcode)
	fmt.Println(a.state)
}
