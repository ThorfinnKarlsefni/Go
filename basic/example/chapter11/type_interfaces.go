package chapter11

import (
	"fmt"
	"math"
)

type Cricle struct {
	radius float32
}

func TypeInter() {
	var araeIntf Shaper

	sql1 := new(Square)
	sql1.side = 5

	araeIntf = sql1
	if t, ok := araeIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is %T\n", t)
	}

	if u, ok := araeIntf.(*Cricle); ok {
		fmt.Printf("The type of areaInft is: %T\n", u)
	} else {
		fmt.Println("areaInfo does not contain a variable of type circle")
	}
}

func (ci *Cricle) Area() float32 {
	return ci.radius * ci.radius * math.Pi

}
