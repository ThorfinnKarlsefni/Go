package datatype

import (
	"fmt"
	"runtime"
)

func Bool() {
	bool1 := true

	if bool1 {
		fmt.Printf("This value is true\n")
		fmt.Printf(runtime.GOOS)
	} else {
		fmt.Printf("The value is false\n")
	}
}
