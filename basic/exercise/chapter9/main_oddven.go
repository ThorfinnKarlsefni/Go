package exercise

import (
	"fmt"
	"go/basic/exercise/chapter9/even"
)

func Oddven() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the intger %d even? %v\n", i, even.Even(i))
	}
}
