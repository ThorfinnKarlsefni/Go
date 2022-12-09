package chapter10

import (
	"fmt"
	"runtime"
)

var m runtime.MemStats

func Run() {
	runtime.ReadMemStats(&m)
	fmt.Printf("%d kb\n", m.Alloc/1024)
}
