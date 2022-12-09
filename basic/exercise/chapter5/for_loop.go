package exercise

import (
	"fmt"
)

func For15() {
	for i := 1; i <= 15; i++ {
		fmt.Println(i)
	}
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}
