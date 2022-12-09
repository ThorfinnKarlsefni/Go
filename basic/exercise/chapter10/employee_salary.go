package exercise

import "fmt"

type employee struct {
	name   string
	salary float32
}

func AddSalary() {
	var e employee
	e.salary = 1000000
	e.name = "Cheung"
	e.giveRaise(0.20)
	fmt.Printf("Employee  now makes %f", e.salary)
}

func (this *employee) giveRaise(param float32) {
	this.salary += this.salary * param
}
