package main

import "fmt"

func main() {
	var a int = 5
	sqrfunc(&a)
	fmt.Println(a)
}

func sqrfunc(x *int) {
	*x *= *x

}
