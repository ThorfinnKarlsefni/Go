package exercise

import (
	"fmt"
	"unicode/utf8"
)

var str1 string = "asSASA ddd dsjkdsjs dk"
var str2 string = "asSASA ddd dsjkdsjsこん dk"

func Len() {
	fmt.Printf("%d \n", len(str1))
	fmt.Printf("%d \n", utf8.RuneCountInString(str1))
	fmt.Printf("%d \n", len(str2))
	fmt.Printf("%d", utf8.RuneCountInString(str2))

}
