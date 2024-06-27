package main

import (
	"fmt"
	"time"
)

func main() {
	go printrout()
	time.Sleep(3 * time.Second)
	fmt.Println("Hello World")
	//time.Sleep(2 * time.Second)
	//myfunction()
}

func printrout() {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello World rout")
}
