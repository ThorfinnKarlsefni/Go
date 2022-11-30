package exercise

import "fmt"

func MapFunction() {
	list := []int{0, 1, 2, 3, 4, 5, 6, 7}
	mf := func(i int) int {
		return i * 10
	}

	println()
	fmt.Printf("%v", mapFunc(mf, list))
}

func mapFunc(mf func(int) int, list []int) []int {
	result := make([]int, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}
