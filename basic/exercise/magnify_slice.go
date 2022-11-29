package exercise

import "fmt"

func MagnifySlice() {
	slice := []int{4, 5, 6}
	factor := [3]int{1, 2, 3}

	m := len(slice)

	newSlice := make([]int, (m)*len(factor))
	copy(newSlice, slice)
	slice = newSlice

	fmt.Printf("slice len is %d, factor len is %d", len(slice), len(factor))

}

var s []int

func TestMagnify() {
	s = []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(s))
	fmt.Println(s)

	s = enlarge(s, 5)
}

func enlarge(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	s = ns
	return s
}
