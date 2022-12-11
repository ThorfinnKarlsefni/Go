package chapter11

import "fmt"

func MapFunctionInterfaceVar() {
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}
		return i
	}

	res1 := mapFunc1(mf, 0, 1, 2, 3, 4, 5)
	for _, v := range res1 {
		fmt.Println(v)
	}

	println()

	res2 := mapFunc1(mf, "0", "1", "2", "3", "4", "5")
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc1(mf func(obj) obj, list ...obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}
