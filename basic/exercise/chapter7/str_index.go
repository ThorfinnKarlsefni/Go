package exercise

import "fmt"

// Segmentation string

func SplitString() {
	str := "life's little bit messy.we all make mistakes."

	for i := 0; i <= len(str); i++ {
		a, b := Split(str, i)
		fmt.Printf("The string %s split at position %d is:%s / %s\n", str, i, a, b)
	}

	fmt.Println(str[:24])
	fmt.Println(str[24:])
}

func Split(s string, pos int) (string, string) {
	return s[:pos], s[pos:]
}
