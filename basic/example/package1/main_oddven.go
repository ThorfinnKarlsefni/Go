package package1

import (
	"fmt"
	"go/basic/example/package1/even"
)

func Oddven() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Is the intger %d even? %v\n", i, even.Even(i))
	}
}
