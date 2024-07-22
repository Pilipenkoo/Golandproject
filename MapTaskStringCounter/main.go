package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Привет, мир! Привет всем. Это пример текста. Пример.Пример.Пример"
	m := make(map[string]int)

	f := func(c rune) bool {
		return c == '!' || c == ' ' || c == ',' || c == '.'
	}

	str_split := strings.FieldsFunc(str, f)
	for _, v := range str_split {
		if _, exists := m[v]; exists {
			m[v] += 1
			continue
		}

		m[v] = 1

	}
	fmt.Println(m)
}
