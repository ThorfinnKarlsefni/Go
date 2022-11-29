package exercise

import "fmt"

func FilterSlice() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s = Filter(s, even)
	fmt.Println(s)
}

// learn closure
func Filter(s []int, fn func(int) bool) []int {
	var p []int

	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}

	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}
