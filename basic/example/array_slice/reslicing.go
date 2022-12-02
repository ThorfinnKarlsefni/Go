package array_slice

import "fmt"

func ReSlice() {
	slice1 := make([]int, 0, 10)

	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
}

func SliceLength() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("arr Length is %d,arr cap is %d\n", len(arr), cap(arr))

	a := arr[7:7]

	fmt.Printf("a Length is %d,a cap is %d", len(a), cap(a))
}
