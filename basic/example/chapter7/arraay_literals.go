package chapter7

import "fmt"

func ArrayLiterals() {
	var arrayKeyValue = [5]string{3: "Chris", 4: "Ron"}

	for i := 0; i < len(arrayKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrayKeyValue[i])
	}
}
