package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("BMI calculator")
	fmt.Println("---------------------------")
	fmt.Println("Enter your Weight in KG: ")
	weight, _ := reader.ReadString('\n')
	weight = strings.Replace(weight, "\n", "", -1)
	fmt.Println("Enter your height in M: ")
	height, _ := reader.ReadString('\n')
	height = strings.Replace(height, "\n", "", -1)
	f_weight, _ := strconv.ParseFloat(weight, 32)
	f_height, _ := strconv.ParseFloat(height, 32)
	bmi := f_weight / (f_height * f_height)
	fmt.Printf("your BMI is %v \n", bmi)

}
