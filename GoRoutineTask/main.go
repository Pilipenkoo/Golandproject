package main

import "fmt"

var a int = 1

func main() {
	done := make(chan bool)
	ch := make(chan bool)
	for i := 0; i < 10; {

		go func() {
			token := <-ch
			for {
				if a >= 1000 {
					done <- true
					ch <- token
					return
				}
				a = a * 2
				ch <- token
			}
		}()
		i++
	}
	ch <- true
	<-done
	fmt.Println("A: ", a)
}
