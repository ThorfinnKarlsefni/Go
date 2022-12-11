package chapter11

import "fmt"

type TopologicalGenus interface {
	Rank() int
}

func (sq *Square) Rank() int {
	return 1
}

func (r Rectangle) Rank() int {
	return 2
}

func MultiInterface() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}
	fmt.Println("Looping through  shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details:", shapes[n])
		fmt.Println("Area of this shape is:", shapes[n].Area())
	}
	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank...")
	for n, _ := range topgen {
		fmt.Println("Shape details:", topgen[n])
		fmt.Println("Topological Genus of this shape is:", topgen[n].Rank())
	}

}
