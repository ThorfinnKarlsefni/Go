package chapter11

import "fmt"

func gI(any interface{}) int {
	if v, ok := any.(Simpler); ok {
		return v.Get()
	}

	return 0
}

func SimpleInterface3() {
	var s Simple = Simple{6}
	fmt.Println(gI(&s))
	var r RSimple = RSimple{60, 60}
	fmt.Println(gI(&r))
}
