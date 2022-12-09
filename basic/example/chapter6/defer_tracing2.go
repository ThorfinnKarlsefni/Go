package chapter6

import "fmt"

func trace1(s string) string {
	fmt.Printf("entering %s\n", s)
	return s
}

func a1() {
	defer un(trace1("a"))
	fmt.Printf("in a\n")
}

func un(s string) {
	fmt.Printf("leaving %s\n", s)
}

func b1() {
	defer un(trace1("b"))
	fmt.Printf("in b\n")
	a1()
}

func Trace2() {
	b1()
}
