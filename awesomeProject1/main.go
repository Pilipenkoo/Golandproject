package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var list1 []int
	list1 = []int{1, 2, 4, 6, 8}
	fmt.Println(Append(list1, 1, 2, 3, 532, 523, 34, 3))
	slice_train()
	reverse(&list1)
	fmt.Println(list1)
	a()

}

func slice_train() {
	slice_list := make([]int, 10, 11)
	slice_list_odd := make([]int, 0, 11)
	var sum_all int = 0
	var avg_all float64
	for i := 0; i < len(slice_list); i++ {
		slice_list[i] = rand.Intn(100)
		sum_all += slice_list[i]
	}
	min, max := slice_list[0], slice_list[0]
	for i := 1; i < len(slice_list); i++ {
		if slice_list[i] < min {
			min = slice_list[i]
		}
		if slice_list[i] > max {
			max = slice_list[i]
		}

	}
	avg_all = float64(sum_all) / float64(len(slice_list))
	for i := 0; i < len(slice_list); i++ {
		if slice_list[i]%2 != 0 {
			slice_list_odd = append(slice_list_odd, slice_list[i])

		}
	}
	fmt.Printf(" Original slice %v\n", slice_list)
	fmt.Printf(" Min:%v\n Max:%v\n", min, max)
	fmt.Printf(" Sum_all %v\n Avg: %v\n Slice without even: %v\n", sum_all, avg_all, slice_list_odd)

}

func handle(list []int) []int {
	newList := make([]int, len(list))
	copy(newList, list)
	fmt.Printf("list now: %v,capacity now: %d\n", newList, cap(newList))
	newList = append(newList, 5)
	fmt.Printf("list after: %v ,capacity after: %d\n", newList, cap(newList))
	return newList
}

func double(numg []int) []int {
	res := make([]int, len(numg))
	for i, num := range numg {
		res[i] = num * 2
	}
	return res
}

func Append(slice []int, elem ...int) []int {
	var res []int
	var resLen = len(slice)
	var resCap = resLen + len(elem) // resCap - результат сложения длины переданного слайса с длиной элементов переданных на добавление

	// Проверка необходимости увеличения вместимости
	if resCap > cap(slice) {
		newCap := resCap
		if newCap < resLen*2 {
			newCap = resLen * 2
		}
		// Создание нового слайса с необходимой длиной и вместимостью
		res = make([]int, resCap, newCap)
		copy(res, slice)
	} else {
		// Расширение существующего слайса до новой длины
		res = slice[:resCap]
	}

	// Копирование новых элементов в слайс
	copy(res[resLen:], elem)
	return res
}
