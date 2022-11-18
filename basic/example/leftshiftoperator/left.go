package leftshiftoperator

import "fmt"

type BitFlag int

const (
	Active BitFlag = 1 << iota
	Send
	Receive
)

func Init() {
	flag := Active | Send

	fmt.Println(flag)
}
