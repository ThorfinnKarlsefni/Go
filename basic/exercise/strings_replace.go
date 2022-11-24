package exercise

import (
	"fmt"
	"strings"
)

func Replace() {
	str := "life's little bit messy?we all make mistakes?"
	res := recursion(str)
	fmt.Println(res)
}

func recursion(str string) string {

	rot13 := func(r rune) bool {
		if r == '?' {
			return true
		}
		return false
	}

	index := strings.IndexFunc(string(str), rot13)

	if index > 0 {
		arr := []byte(str)
		arr[index] = '.'
		recursion(string(arr))
		println(string(arr), "111")
		return string(arr)
	}
	return str
}
