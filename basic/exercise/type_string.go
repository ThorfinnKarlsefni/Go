package exercise

import (
	"fmt"
)

type T struct {
	a int
	b float32
	c string
}

func TypeString() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t)
}

func (t *T) String() string {

	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}
