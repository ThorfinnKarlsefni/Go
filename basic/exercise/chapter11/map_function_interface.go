package chapter11

import "fmt"

type obj interface{}

func MapFuntionInterface() {
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		}

		return i
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, isl)

	for _, v := range res1 {
		fmt.Println(v)
	}

	println()

	ssl := []obj{"life's little bit messy", "we all make mistakes"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Println(v)
	}
}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))
	for ix, v := range list {
		result[ix] = mf(v)
	}
	// Equivalent

	//for ix := 0; ix < len(list); i++ {
	//	result[ix] = mf(list)
	//}

	return result
}
