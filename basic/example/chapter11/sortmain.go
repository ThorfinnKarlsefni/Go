package chapter11

import (
	"fmt"
	"go/basic/example/chapter11/sort"
)

func Ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}

	fmt.Printf("The sorted array is: %v\n", a)
}
