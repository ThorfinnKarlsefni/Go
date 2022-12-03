package map1

import (
	"fmt"
	"sort"
)

func SortMap() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key %v,Value: %v /", k, v)
	}

	keys := make([]string, len(barVal))

	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v,Value: %v/", k, barVal[k])
	}
}
