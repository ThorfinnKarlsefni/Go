package struct_func

func (n *NamedPoint) Abs() float64 {
	return n.Point1.Abs() * 100
}
