package main

import (
	"fmt"
)

func main() {
	// Создаем два канала
	intChannel := make(chan int)
	doubleChannel := make(chan int)

	go doubleInt(intChannel, doubleChannel)

	go sendInt(intChannel)

	for res := range doubleChannel {
		fmt.Println(res)
	}
}

func sendInt(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func doubleInt(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
	close(out)
}
