package main

import (
	"fmt"
	"sync"
	"time"
)

func wg_train() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i + 1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Done")
}

func sequential() {
	for i := 0; i < 10; i++ {
		fmt.Println(i + 1)
	}
	fmt.Println("Done")
}

func main() {
	start := time.Now()
	wg_train()
	duration := time.Since(start)
	fmt.Printf("Time taken with goroutines: %v\n", duration)

	start = time.Now()
	sequential()
	duration = time.Since(start)
	fmt.Printf("Time taken sequentially: %v\n", duration)
}
