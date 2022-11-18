package alias

import "fmt"

type TZ int

func TypeAlias() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c has the value : %d", c)
}
