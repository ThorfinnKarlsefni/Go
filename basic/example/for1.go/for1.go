package for1

import "fmt"

func Init() {
	for i := 0; i < 10; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
