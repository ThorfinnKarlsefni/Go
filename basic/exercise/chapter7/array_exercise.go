package exercise

import (
	"fmt"
	"time"
)

func ArrayValue() {
	var arr1 [5]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	arr2 := arr1
	arr2[2] = 888

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Arrary arr1 at index %d is %d\n", i, arr1[i])
	}

	fmt.Println()
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("Arrary arr1 at index %d is %d\n", i, arr2[i])
	}
}

func ArrayFor() {
	arr := [16]int{}
	for i := 0; i < 16; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}

func FibonacciFor() {
	start := time.Now()

	arr := [51]int{}
	for i := 0; i < 50; i++ {
		if i < 2 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
		fmt.Println(arr[i])
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Println(delta)
}
