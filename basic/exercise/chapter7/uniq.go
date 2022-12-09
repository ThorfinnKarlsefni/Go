package exercise

import "fmt"

func Uniq() {
	arr := []byte{'a', 'b', 'a', 'a', 'a', 'c', 'd', 'e', 'f', 'g'}
	arru := make([]byte, len(arr))

	ixu := 0
	tmp := byte(0)

	for _, val := range arr {

		if val != tmp {
			arru[ixu] = val
			fmt.Printf("%c", arru[ixu])
			ixu++
		}
		tmp = val
	}
}
