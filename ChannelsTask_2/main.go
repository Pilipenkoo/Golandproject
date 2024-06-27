package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	chint1 := make(chan int)
	chint2 := make(chan int)
	sumAndSend := make(chan int)
	multiply := make(chan int)

	// Отправка случайных чисел в два канала
	go func() {
		for i := 0; i < 5; i++ {
			chint1 <- rand.Intn(10)
			chint2 <- rand.Intn(15)
		}
		close(chint1)
		close(chint2)
	}()

	// Суммирование чисел и отправка результатов в следующий канал
	go sendSum(chint1, chint2, sumAndSend)

	// Умножение суммы на 2 и отправка результатов в следующий канал
	go multiplyInt(sumAndSend, multiply)

	// Получение и вывод результатов
	for result := range multiply {
		fmt.Println(result)
	}
}

func sendSum(chint1, chint2 <-chan int, out chan<- int) {
	for {
		a, ok1 := <-chint1
		b, ok2 := <-chint2
		if !ok1 || !ok2 {
			close(out)
			return
		}
		sum := a + b
		out <- sum
	}
}

func multiplyInt(in <-chan int, out chan<- int) {
	for val := range in {
		out <- val * 2
	}
	close(out)
}
