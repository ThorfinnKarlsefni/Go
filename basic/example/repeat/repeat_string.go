package repeat

import (
	"fmt"
	"strings"
)

func Init() {
	var origS string = "Hi there!"
	var newS string

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}
