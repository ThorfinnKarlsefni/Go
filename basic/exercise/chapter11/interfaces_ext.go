package chapter11

import "fmt"

type Triangle struct {
	base, height float32
}

func (t Triangle) Area() float32 {
	return t.base * t.height * 0.5
}

type AreaInterface interface {
	Area() float32
}

func fi2(a AreaInterface) float32 {
	return a.Area()
}

type PeriInterface interface {
	Perimeter(int) int
}

func (s Square) Perimeter(i int) int {
	return i + i
}

func InterfaceExt() {
	t := Triangle{10, 6}
	//var t Triangle
	//t.bottom = 10
	//t.height = 5
	fmt.Println(fi2(t))
	s := Square{}
	var p1 PeriInterface
	p1 = s
	fmt.Println(p1.Perimeter(10))
}
