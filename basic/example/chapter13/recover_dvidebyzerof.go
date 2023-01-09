package chapter13

import "fmt"

func badCall1() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}

func test1() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panicing %s\r\n", e)
		}
	}()

	badCall1()
	fmt.Printf("After bad call\r\n")
}

//func DivisionByZero() {
//	fmt.Println("calling test\r\n")
//	test1()
//	fmt.Println("test completed\r\n")
//}
