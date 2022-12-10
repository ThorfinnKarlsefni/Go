package chapter11

import "fmt"

type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	id int
}

func (p *Simple) Get() int {
	return p.id
}

func (p *Simple) Set(id int) {
	p.id = id
}

func fI(it Simpler) int {
	it.Set(5)
	return it.Get()
}

func SimpleInterface() {
	s := new(Simple)
	// var s Simple
	// fmt.Println(s)
	fmt.Println(fI(s))
}
