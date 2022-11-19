package exercise

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FizzBuzz = 15
)

func Fizz_Buzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println(FizzBuzz)
		case i%3 == 0:
			fmt.Println(FIZZ)
		case i%5 == 0:
			fmt.Println(BUZZ)
		default:
			fmt.Println(i)
		}
	}
}
