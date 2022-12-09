package exercise

import "fmt"

type A struct {
	area float32
	B
}
type B struct {
	in1 int
	in2 string
}

func Anonymous_struct() {
	square := A{}
	square.area = 12.3
	square.in1 = 19
	square.in2 = "we all make mistakes"
	fmt.Println(square)
}
