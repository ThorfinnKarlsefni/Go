package chapter4

import (
	"fmt"
	"os"
)

func Var() {
	// var identifier [type] = value
	var a int = 15
	var i = 5
	var b bool = false
	var str string = "Go says hello to the world"

	// var (
	// 	a   int = 15
	// 	b       = 6
	// 	str     = "此刻我在异乡的夜里 想念着你越来越远"
	//	numShip = 56
	//	city string
	// )
	var (
		HOME   = os.Getenv("HOME")
		USER   = os.Getenv("USER")
		GOROOT = os.Getenv("GOROOT")
	)

	fmt.Println(a, i, b, str)
	fmt.Println(HOME, USER, GOROOT)
}
