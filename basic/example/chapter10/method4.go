package chapter10

func (n *NamedPoint) Abs() float64 {
	return n.Point1.Abs() * 100
}
