package main

import "fmt"

func longestPalindrome(s string) string {
	var originalString string
	originalString = s
	var reverse = []rune(originalString)
	var result string
	originalrunes := []rune(originalString)
	predel := originalrunes
	for i := 0; i < len(predel)/2; i++ {
		reverse[i] = originalrunes[len(originalrunes)-i]
	}
	for i := 0; i < len(originalrunes)/2; i++ {
		if string(originalrunes) == string(reverse) {
			result = string(originalrunes)
			return result
		}

	}

	return result
}

func main() {
	x := longestPalindrome("aab")
	fmt.Printf(x)
}
