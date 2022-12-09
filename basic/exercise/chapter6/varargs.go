package exercise

import "fmt"

func VarArgs() {
	length(1, 3, 5, 7, 8, 9, 11)

}

func length(s ...int) {
	for _, v := range s {
		fmt.Printf("value is %d\n", v)
	}
}
