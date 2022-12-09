package exercise

import "fmt"

func SumArrayExercise() {
	arrF := []float64{78.11, 33.33, 11.2331, 113}
	arrSum := SumArr(arrF)
	sum, avg := SumAndAvg(arrF)
	fmt.Println(arrSum)
	fmt.Println(sum, avg)
}

func SumArr(arr []float64) (sum float64) {
	if len(arr) > 0 {
		for _, v := range arr {
			sum += v
		}
	} else {
		sum = 0
	}
	return
}

func SumAndAvg(arr []float64) (sum, avg float64) {
	sum = SumArr(arr)

	if len(arr) != 0 {
		avg = sum / float64(len(arr))
	}

	return
}
