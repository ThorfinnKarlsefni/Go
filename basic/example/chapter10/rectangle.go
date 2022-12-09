package chapter10

import (
	"fmt"
)

type Rectangle struct {
	len, wid int
}

func (r *Rectangle) Area() int {
	return r.len * r.wid
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.len + r.wid)
}

func Rec() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is:", r1)
	fmt.Println("Rectangle area is:", r1.Area())
	fmt.Println("Rectangle Perimeter is:", r1.Perimeter())

}
