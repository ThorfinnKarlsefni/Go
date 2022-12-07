package exercise

import "fmt"

type inter interface {
	Id()
	setId()
}

type Base struct {
	id int
	inter
}

func (this *Base) Id() int {
	return this.id
}

func (this *Base) setId(id int) {
	this.id = id
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float32
}

func Employee1() {
	emp := new(Employee)

	emp.FirstName = "Cheung"
	emp.LastName = "ziming"
	emp.salary = 300000
	fmt.Println(emp.Id())
	emp.setId(0007)
	fmt.Println(emp.Id())
}
