package exercise

import (
	"fmt"
	"time"
)

var result [15]int

func ExampleFibonacci() {
	start := time.Now()
	res := fibarray(15)
	for ix, fib := range res {
		fmt.Printf("The %d-th Fibonacci number is %d\n", ix, fib)
	}

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longcalculator took of time %s\n", delta)
}

func fibarray(term int) []int {
	farr := make([]int, term)
	farr[0], farr[1] = 1, 1

	for i := 2; i < term; i++ {
		farr[i] = farr[i-1] + farr[i-2]
	}

	return farr
}

func FibonacciArrayNotCache() {

	start := time.Now()
	slice1 := make([]int, 15)
	for i := 0; i < len(slice1); i++ {

		if i <= 1 {
			slice1[i] = 1
		} else {
			slice1[i] = slice1[i-1] + slice1[i-2]
		}

		fmt.Println(slice1[i])
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("LongCalcultaion took of time %s\n ", delta)
}

func FibonacciTestRecursion() {
	start := time.Now()
	for i := 0; i < 15; i++ {
		res := fibonacciClosure(i)
		fmt.Printf("fibonacci(%d) is :%d\n", i, res)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longcalculator took of time %s\n", delta)

	//fmt.Printf("LongCalculation took of time %s\n",del)
}

func FibonacciTestRecursionCache() {
	start := time.Now()
	for i := 0; i < 15; i++ {
		//result := fibonacciClosure(i)
		res := fibonacciSliceCalculator(i)
		fmt.Printf("fibonacci(%d) is :%d\n", i, res)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longcalculator took of time %s\n", delta)

	//fmt.Printf("LongCalculation took of time %s\n",del)
}

func fibonacciClosure(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciClosure(n-1) + fibonacciClosure(n-2)
	}
	return
}

func fibonacciSliceCalculator(i int) (res int) {
	if result[i] != 0 {
		res = result[i]
		return
	}

	if i <= 1 {
		res = 1
	} else {
		res = fibonacciSliceCalculator(i-1) + fibonacciSliceCalculator(i-2)
	}
	result[i] = res
	return
}
