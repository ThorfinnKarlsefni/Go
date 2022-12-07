package struct_func

import "fmt"

type Customer1 struct {
	Name string
	Log1
}

type Log1 struct {
	msg string
}

func Embed1() {
	c := new(Customer1)
	c.Name = "Cheung"
	c.Log1 = Log1{"心里想有一些话 我们先不讲"}
	c.Add("你得到你想要的吗 换来的确实铁石心肠")
	fmt.Println(c)
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log1) String() string {
	return l.msg
}

func (c *Customer1) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log1)
}
