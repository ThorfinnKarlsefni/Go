package for1

import "fmt"

func Goto2() {
	a := 1
	goto TARGET
	b := 9
TARGET:
	b += a
	fmt.Printf("a is %v *** b is %v", a, b)
}
