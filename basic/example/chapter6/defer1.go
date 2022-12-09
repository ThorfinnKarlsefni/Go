package chapter6

import "fmt"

func Defer1() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
