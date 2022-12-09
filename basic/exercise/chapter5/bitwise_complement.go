package exercise

import "fmt"

func Bitwise() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("the complement of %d is : %b\n", i, ^i)
	}
}
