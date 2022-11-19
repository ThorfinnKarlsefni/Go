package goos

import (
	"fmt"
	"os"
	"runtime"
)

func Init() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}