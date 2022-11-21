package deferandtrace

import "fmt"

func Init() {
	function1()
}

func function1() {
	fmt.Printf("In function1 at the top!\n")
	defer function2()
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2:Deferred untill the end of the calling funciton!")
}
