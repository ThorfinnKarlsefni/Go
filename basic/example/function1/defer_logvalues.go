package function1

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d,%v", s, n, err)
	}()
	return 7, io.EOF
}

func Log1() {
	func1("GO")
}
