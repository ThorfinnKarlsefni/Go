package package1

import (
	"fmt"
	"unsafe"
)

func Unsafe() {

	var i int = 10
	size := unsafe.Sizeof(i)

	fmt.Println("The size of an int is:", size)
}
