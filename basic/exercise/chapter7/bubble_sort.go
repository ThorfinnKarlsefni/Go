package exercise

import "fmt"

func BubbleSort() {
	n := []int{1, 2, 453, 123, 101, 3, 11, 456, 78, 9, 23}

	for i := 0; i < len(n); i++ {
		for j := 1; j < i; j++ {
			if n[i] < n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}

	fmt.Println(n)
}
