package chapter10

import "fmt"

type B1 struct {
	thing int
}

func (b *B1) change() { b.thing = 1 }

func (b B1) write() string { return fmt.Sprint(b) }

func PointerValue() {
	var b1 B1
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B1)
	b2.change()
	fmt.Println(b2.write())
}
