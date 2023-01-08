package chapter13

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	badCall()
	fmt.Printf("After bad Call\r\n")
}

func Pan() {
	fmt.Println("Calling test\r\n")
	test()
	fmt.Println("Test completed\r\n")
}
