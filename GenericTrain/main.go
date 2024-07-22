package main

import "fmt"

func generic[T any](x T) {
	fmt.Println(x)
}

func main() {
	generic(14)
	generic("dwadugaw")
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	var s string = "dwagwag"
	fmt.Println(string(s[5]))
	maps_train()

}

func lengthOfLongestSubstring(s string) int {
	// Инициализация мапы для хранения последних позиций символов
	charIndexMap := make(map[byte]int)
	maxLength := 0
	start := 0

	// Проход по строке
	for end := 0; end < len(s); end++ {
		// Если символ уже встречался и его индекс в пределах текущего окна
		if index, found := charIndexMap[s[end]]; found && index >= start {
			fmt.Println(index)
			// Перемещение начала окна на позицию после последнего появления символа
			start = index + 1
		}

		// Обновление мапы последней позиции текущего символа
		charIndexMap[s[end]] = end

		// Вычисление текущей длины подстроки
		currentLength := end - start + 1

		// Обновление максимальной длины подстроки
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func lenght2(s string) int {
	charIndexMap := make(map[byte]int)
	start := 0
	maxLength := 0
	currentlength := 0

	for end := 0; end <= len(s); end++ {
		if index, found := charIndexMap[s[end]]; found && index >= start {
			start = index + 1
		}
		charIndexMap[s[end]] = end
		currentlength = end - start + 1
		if currentlength > maxLength {
			maxLength = currentlength
		}
	}
	return maxLength

}
