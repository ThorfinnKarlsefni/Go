package exercise

import "fmt"

func StringAppend() {

	b := []byte{213, 21, 21, 1, 2, 74, 8}
	var s string = "maweiwei"
	b = append(b, s...)
	fmt.Println(b)
}

func Convent() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
}
