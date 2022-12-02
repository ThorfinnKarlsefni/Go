package datatype

import "fmt"

type BitFlag int

const (
	Active BitFlag = 1 << iota
	Send
	Receive
)

func Left() {
	flag := Active | Send

	fmt.Println(flag)
	fmt.Println(Active)
	fmt.Println(Send)
}
