package exercise

import (
	"fmt"
)

func Fact() {
	result := factorial1(4)
	fmt.Println(result)
}

// func factorial(ret int) (result int) {
// 	if ret > 0 {
// 		result = ret * factorial(ret-1)
// 		return result
// 	} else {
// 		result = 1
// 		return result
// 	}
// }

func factorial1(ret int) int {
	var result1 int
	if ret > 0 {
		result1 = ret * factorial1(ret-1)
	} else {
		result1 = 1
	}

	return result1
}
