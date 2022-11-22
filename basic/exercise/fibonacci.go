package exercise

import "fmt"

func Fibonacci() {
	res := fibo(2)
	fmt.Printf("fibonacci is value %d", res)
}

func fibo(i int) (res int) {
	if i <= 0 {
		res = 1
	} else {
		res = fibo(i-1) + fibo(i-2)
	}

	return
}
