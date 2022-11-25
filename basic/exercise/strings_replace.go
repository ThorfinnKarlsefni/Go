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
	arr := []byte(str)

	rot13 := func(r rune) bool {
		if r == '?' {
			return true
		}
		return false
	}

	index := strings.IndexFunc(string(arr), rot13)

	if index > 0 {
		arr[index] = '.'
		recursion(string(arr))
	}
	return string(arr)
}
