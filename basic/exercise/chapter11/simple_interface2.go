package chapter11

type RSimple struct {
	id int
	j  int
}

func (p *RSimple) Get() int {
	return p.j
}

func (p *RSimple) Set(u int) {
	p.j = u
}

func fI1(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(5)
		return it.Get()
	case *RSimple:
		it.Set(50)
		return it.Get()
	default:
		return 99
	}
	return 0
}

func Interface2() {
	var a Simple
	println(fI1(&a))
	var b RSimple
	println(fI1(&b))
}
