package package1

import (
	"fmt"
	"go/basic/example/package1/fibo"
)

var nextFibo int
var op string

func Fibo() {
	op = "+"
	calls()
	fmt.Printf("Change of operation from + to *\n")
	nextFibo = 0
	op = "*"
	calls()
}

func calls() {
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
	fmt.Println("...")
	next()
}

func next() {
	result := 0
	nextFibo++
	result = fibo.Fibonacci(op, nextFibo)
	fmt.Printf("fibonacci(%d) is: %d\n", nextFibo, result)
}
