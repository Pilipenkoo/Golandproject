package main

import "fmt"

func eng_rus() {
	englishLowercase := []rune("abcdefghijklmnopqrstuvwxyz ")
	russianLowercase := []rune("абвгдеёжзийклмнопрстуфхцчшщъыьэюя")
	for i, v := range englishLowercase {
		j := i + 1
		if j%3 == 0 && j%5 == 0 && i != 0 {
			fmt.Printf("Привет\n")
		} else if j%3 == 0 && i != 0 {
			fmt.Printf("%v\n", j)
		} else if j%5 == 0 && i != 0 {
			fmt.Printf("%s\n", string(russianLowercase[i]))
		} else {
			fmt.Printf("%s\n", string(v))
		}
	}

}

/* Условие задачи:
вывести английский алфавит со следующими условиями:
	1. каждая 3 буква выводится её порядковый номер
	2. каждая 5 выводится русская буква, с соответствующим порядковым номером
	3. 15 буква выводится слово Привет*/
