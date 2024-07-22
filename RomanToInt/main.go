package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MDCXCV"))

}
func romanToInt(s string) int {
	count := 0
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	srunes := []rune(s)
	if len(srunes) == 1 {
		count = m[srunes[0]]
		return count
	}
	for i := 0; i < len(srunes)-1; i++ {
		if m[srunes[i]] >= m[srunes[i+1]] {
			count += m[srunes[i]]
			if i+1 == len(srunes)-1 {
				count += m[srunes[i+1]]
			}
		} else {
			for j := i + 1; j < len(srunes); j++ {
				result := m[srunes[j]] - m[srunes[i]]
				count += result
				i = j
				break
			}
			if i+1 == len(srunes)-1 {
				count += m[srunes[i+1]]
			}
		}

	}
	return count
}
