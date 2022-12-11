package chapter11

import "fmt"

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func NodeStructures() {
	root := NewNode(nil, nil)
	root.SetData("root Node")
	a := NewNode(nil, nil)
	a.SetData("left Node")
	b := NewNode(nil, nil)
	b.SetData("right Node")
	root.le = a
	root.ri = b
	fmt.Printf("%v\n", root)
}
