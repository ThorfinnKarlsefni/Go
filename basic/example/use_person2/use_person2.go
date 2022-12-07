package use_person2

import (
	"fmt"
	"go/basic/example/person2"
)

func Person2() {
	p := new(person2.Person)

	p.SetFirstName("Askeladd")
	fmt.Println(p.FirstName())
}
