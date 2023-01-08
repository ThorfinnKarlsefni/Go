package chapter12

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5() {
	hasher := md5.New()
	b := []byte{}

	io.WriteString(hasher, "test")
	fmt.Printf("result %x\n", hasher.Sum(b))
	fmt.Printf("result %d\n", hasher.Sum(b))
}
