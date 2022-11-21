package function1

import "fmt"

func Multiply(a, b int, reply *int) {
	*reply = a + b
}

func Side() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply) // Multiply: 50
}
