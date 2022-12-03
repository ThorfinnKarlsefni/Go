package map1

import "fmt"

func MapFunc() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println(mf)
}
