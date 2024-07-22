package main

import "fmt"

func main() {
	//slice := make([]int, 10, 20)
	x := 141
	fmt.Println(isPalindrome(x))

}
func isPalindrome(x int) bool {
	original_number := x
	reverse_number := 0
	if x < 0 {
		return false
	}
	for x > 0 {
		digit := x % 10
		reverse_number = reverse_number*10 + digit
		x = x / 10
	}
	return original_number == reverse_number
}
