package main

import "fmt"

func reverse(slice *[]int) {
	s := *slice
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

}

func a() {
	x := []int{}
	x = append(x, 0) // cap: 1
	x = append(x, 1) // cap: 2
	x = append(x, 2) // 0, 1, 2 cap: 4, len 3 [0, 1, 2, 4]

	y := append(x, 3) // y: 0, 1, 2, 4
	z := append(x, 4) //

	fmt.Println(y, z) // ответ: 0, 1, 2, 4      0, 1, 2, 4
	fmt.Println("x:", x)
}
