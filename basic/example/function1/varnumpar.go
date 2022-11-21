package function1

import "fmt"

func Var() {
	x := min(1, 2, 3, 4)
	fmt.Printf("The minimum is %d\n", x)
	slice := []int{7, 8, 2, 1, 4}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is %d", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}
