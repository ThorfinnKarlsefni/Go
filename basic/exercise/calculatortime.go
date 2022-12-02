package exercise

import (
	"fmt"
	rec "go/basic/example/array_slice"
	"time"
)

func TestCalculationTime() {
	start := time.Now()
	rec.Init()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("lognCalculation took this amount of time:%s\n", delta)
}
