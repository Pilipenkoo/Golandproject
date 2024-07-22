package main

import "fmt"

func maps_train() {
	strmap := make(map[string]int, 10)

	strmap["one"] = 14
	fmt.Println(strmap["one"])
	strmap["two"] = 12
	fmt.Println(strmap["two"])
	strmap["three"] = 13
	fmt.Println(strmap["three"])
	strmap["one"] = 2141
	fmt.Println(strmap)
	fmt.Println(len(strmap))
}
