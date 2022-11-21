package function1

import "fmt"

func Indertifier() {
	i1, _, i2 := threeValue()
	fmt.Printf("The int: %d,the float %f \n", i1, i2)
}

func threeValue() (int, int, float32) {
	return 5, 6, 99.00
}
