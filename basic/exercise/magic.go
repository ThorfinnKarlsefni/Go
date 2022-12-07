package exercise

import "fmt"

type Base1 struct{}

func (Base1) Magic() {
	fmt.Println("base magic") // 1 // 3 4
}

func (self Base1) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base1
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic") // 2
}

func Magic() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
