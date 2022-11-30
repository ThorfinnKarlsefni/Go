package exercise

import "fmt"

func Reverse() {
	str := "Google"
	//fmt.Println(slice)
	str2 := split(str)
	fmt.Println("The string %s transformed is %s\n", str, str2)

}

func split(s string) string {
	mid := len(s) / 2
	return s[mid:] + s[:mid]
}
