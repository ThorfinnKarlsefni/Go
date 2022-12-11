package chapter11

import "fmt"

func ints() {
	data := InArray{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	m := Min(data)
	fmt.Printf("the minimum of the array is: %v\n", m)
}

func strings() {
	data := StringArray{"ddd", "eee", "bbb", "ccc", "aaa"}
	m := Min(data)
	fmt.Printf("The minnimum of the array is :%v\n", m)
}

func MinMain() {
	ints()
	strings()
}
