package chapter10

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float64
}

func Embedd() {
	b := B{A{1, 2}, 3.0, 3}
	fmt.Println(b.A)
	fmt.Println(b.ax, b.ay, b.bx, b.by)
}
