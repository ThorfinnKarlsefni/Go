package chapter8

import (
	"fmt"
)

func MapSlice() {
	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}

	fmt.Printf("Version A: Value of items: %v\n", items)
	fmt.Println(items[1][1])

	// 在 B 版本中获得的项只是 map 值的一个拷贝而已
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)

		item[1] = 2
	}
	//fmt.Println(items2[])
	fmt.Printf("Version B: Value of items: %v\n", items2)
}
