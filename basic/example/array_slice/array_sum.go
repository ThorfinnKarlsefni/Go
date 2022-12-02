package array_slice

import "fmt"

func ArraySum() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)

	fmt.Printf("The sum of the array is:%f", x)
}

func Sum(arr *[3]float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}
