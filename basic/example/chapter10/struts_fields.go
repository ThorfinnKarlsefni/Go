package chapter10

import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func Example1() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 10.33
	ms.str = "we all make mistakes"

	fmt.Printf("the int is %d\n", ms.i1)
	fmt.Printf("the float is %f\n", ms.f1)
	fmt.Printf("the str is %s\n", ms.str)
	fmt.Println(ms)
}
