package main

import "fmt"

func main() {
	mp := make(map[string]float32, 5)
	sum := 0.0
	avg := 0.0
	mp["Ain"] = 4.00
	mp["Mike"] = 5.00
	mp["Ivan"] = 4.00
	for _, v := range mp {
		sum += float64(v)
	}
	avg = sum / float64(len(mp))
	fmt.Printf("Average value: %f\n", avg)
}
