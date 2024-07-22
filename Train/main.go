package main

import (
	"fmt"
	"time"
)

func main() {
	timeStart := time.Now()

	ch1 := worker()
	ch2 := worker()

	<-ch1
	<-ch2

	fmt.Println(int(time.Since(timeStart).Seconds())) // Ожидаем, что будет 3 секунды
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
