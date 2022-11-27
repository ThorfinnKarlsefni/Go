package exercise

import "fmt"

func MinSlice() {
	slice := []int{1, 3, 5, 11, 33, 14, 213, 466, 13, 132}
	min := MinS(slice)
	max := MaxS(slice)
	fmt.Println(min)
	fmt.Println(max)

}

func MinS(res []int) (min int) {
	min = res[0]
	for i := 1; i < len(res); i++ {
		if min > res[i] {
			min = res[i]
		}
	}
	return
}

func MaxS(res []int) (max int) {
	max = res[0]
	for i := 1; i < len(res); i++ {
		if max < res[i] {
			max = res[i]
		}
	}

	return
}
