package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	nums := []int{1, 2, 3, 4, 5}
	target := 6
	fmt.Println(twoSum(nums, target))
	go one(ch)

	go Two(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

func one(ch chan int) {
	time.Sleep(5 * time.Second)
	fmt.Println("One func")
	ch <- 5
}

func Two(ch chan int) {
	time.Sleep(55 * time.Millisecond)
	fmt.Println("Two func")
	ch <- 7
}
