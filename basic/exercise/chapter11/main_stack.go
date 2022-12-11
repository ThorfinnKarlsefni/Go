package chapter11

import (
	"fmt"
	"go/basic/exercise/chapter11/stack"
)

var st1 stack.Stack

func Stack() {
	st1.Push("Brown")
	st1.Push("Cheung_Ne123")
	st1.Push("9451")
	st1.Push([]string{"Java", "C++", "Python", "C#", "Ruby"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	//fmt.Println(st1)
}
