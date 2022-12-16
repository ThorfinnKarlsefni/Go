package chapter12

import (
	"fmt"
	"os"
	"strings"
)

func OsArgs(){
	who := "Alice"
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:],"")
	}

	fmt.Println("Good Morning",who)
}