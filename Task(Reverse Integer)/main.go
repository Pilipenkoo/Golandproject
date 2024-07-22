package main

import (
	"fmt"
)

func reverse(x int) int {
	defaultnumber := 0
	originalNumber := x
	number := originalNumber
	for number != 0 {
		digit := number % 10
		defaultnumber = defaultnumber*10 + digit
		number = number / 10
	}
	if defaultnumber < 2147483648 && defaultnumber > -2147483648 {
		return defaultnumber
	} else {
		return 0
	}
}

func main() {
	var x int64 = -120
	fmt.Println(reverse(x))
}
