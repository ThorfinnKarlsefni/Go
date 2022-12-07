package struct_func

import (
	"fmt"
	"math"
)

type Point1 struct {
	x, y float64
}

func (p *Point1) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point1
	name string
}

func Method3() {
	n := &NamedPoint{Point1{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}
