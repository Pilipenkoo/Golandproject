package main

import "fmt"

func main() {
	var (
		first  int = 2
		second int = 5
	)
	fmt.Println(calculate(first, second, Sum))
	m := make(map[string]int, 10)
	m["One"] = 1
	eng_rus()
}

func Sum(a int, b int) int {
	return a + b
}
func calculate(x, y int, abc func(x, y int) int) int {
	return abc(x, y)
}
