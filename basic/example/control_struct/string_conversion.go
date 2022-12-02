package control_struct

import (
	"fmt"
	"strconv"
)

func Str() {
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of this is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is :%d\n", an)

	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
