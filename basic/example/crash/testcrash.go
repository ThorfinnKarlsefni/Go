package crash

import "fmt"

func TestCrash() {
	// a := 1231
	var p *int

	fmt.Println(*p)
}
