package main

import (
	"fmt"
)

func chread(a <-chan int, b chan<- bool) {
	fmt.Printf("%d", <-a)
	b <- true

}

func chwrite(a chan<- int, b int) {

	a <- b

}

func chswrite(a chan<- rune) {
	str := "ABCDEFGHIJKLMN"
	runes := []rune(str)
	for i := 0; i < len(str); i++ {
		a <- runes[i]
	}
	close(a)
}

func chsread(a <-chan rune, b chan<- bool) {
	for r := range a {
		fmt.Printf("%c\n", r)
	}
	b <- true
}

func ch3write(a chan<- int) {
	for i := 0; i < 3; i++ {
		a <- i
	}
	close(a)
}

func ch3read(a <-chan int, b chan<- bool) {
	for r := range a {
		fmt.Printf("%d\n", r)
	}
	b <- true
}
func worker(done chan<- struct{}) {
	// Выполнение некоторой работы (например, вывод сообщения)
	fmt.Println("Горутина выполняет работу")


	// Отправляем сигнал завершения через канал
	done <- struct{}{}
}

func main() {
	// Создаем канал для синхронизации
	done := make(chan struct{})

	// Запускаем горутину
	go worker(done)

	// Ожидаем сигнала завершения от горутины
	<-done

	// Выводим сообщение о завершении работы
	fmt.Println("Работа горутины завершена")
}
