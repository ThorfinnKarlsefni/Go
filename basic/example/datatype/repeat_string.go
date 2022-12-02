package datatype

import (
	"fmt"
	"strings"
)

func Repeat() {
	var origS string = "Hi there!"
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}
