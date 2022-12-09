package chapter9

import "fmt"

func PackTest() {
	var test1 string
	test1 = ReturnStr()

	fmt.Printf("ReturnStr from package1: %s\n", test1)
	fmt.Printf("Integer form package1: %d\n", Pack1Int)
}
