package chapter10

import (
	"fmt"
)

func Person2() {
	p := new(Person)

	p.SetFirstName("Askeladd")
	fmt.Println(p.FirstName())
}
