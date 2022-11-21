package function1

import "fmt"

func Compare() {
	MinMax(5, 6)
}

func MinMax(a int, b int) {
	var max int
	var min int

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	fmt.Printf("max : %d, min : %d", max, min)
}
