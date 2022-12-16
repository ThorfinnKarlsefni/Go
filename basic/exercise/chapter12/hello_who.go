package chapter12

import (
	"fmt"
	"os"
	"strings"
)

func Hello_Who()  {
	hello := "Hello "
	if len(os.Args) > 1 {
		hello += strings.Join(os.Args[1:]," ")
	}
	fmt.Println(hello)
}