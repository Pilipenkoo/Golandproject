package main

func main() {
	a := []int{8, 3, 1, 5, 2, 7, 4, 6}
	n := len(a)
	for i := range n {
		sample := false
		for j := range n - i - 1 {

			if a[j] > a[j+1] {
				a[j+1], a[j] = a[j], a[j+1]
				sample = true
			}

		}
		if !sample {
			break
		}

	}
}
