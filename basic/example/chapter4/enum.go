package chapter4

import "fmt"

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func Enum() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(1 << 4)
	fmt.Println(1 << 10)
}
