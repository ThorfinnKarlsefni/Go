package map1

import (
	"fmt"
	"sort"
)

func Drink() {
	drinks := map[string]string{
		"Coca cola": "可口可乐",
		"Pepsi":     "百事可乐",
		"Hot pot":   "火锅",
		"Juice":     "果汁",
		"Ice":       "冰块",
	}

	sortMap := make([]string, len(drinks))
	//sortMap := [5]string{}
	i := 0
	for k, _ := range drinks {
		sortMap[i] = k
		i++
	}
	sort.Strings(sortMap)
	for k, v := range sortMap {
		fmt.Printf("value is %d,drinks is %v\n", k, drinks[v])
	}
}
