package chapter4

import (
	"fmt"
	"strings"
)

func Index() {
	var str string = "Hi,I'm Thorfinn,Hi."

	fmt.Printf("The position of \"Thorfinn\" is : ")
	fmt.Printf("%d\n", strings.Index(str, "Thorfinn"))

	fmt.Printf("The position of the first instance of \"Hi\" is:")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is:")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is:")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}
