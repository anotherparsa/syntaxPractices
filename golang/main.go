package main

import (
	"fmt"
)

func main() {

	twodimentionalarray := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(twodimentionalarray[0][2]) // 3
	fmt.Println(twodimentionalarray[1][1]) // 5

}
