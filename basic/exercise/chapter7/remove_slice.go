package exercise

import "fmt"

func RemoveSliceString() {
	s := []string{"h", "e", "l", "l", "l", "l", "o"}
	str := removeSlice(s, 2, 4)
	fmt.Println(str)
}

func removeSlice(str []string, start int, end int) []string {
	n := len(str[:start]) + len(str[end:])
	newSlice := make([]string, n)
	at := copy(newSlice, str[:start])

	copy(newSlice[at:], str[end:])
	str = newSlice
	return str
}
