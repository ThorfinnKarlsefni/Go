package reslice

import "fmt"

func SliceAppend() {
	sl_form := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_form)
	fmt.Println(n)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
