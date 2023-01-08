package chapter13

import (
	"errors"
	"fmt"
	"math"
)

var errNotFound error = errors.New("not found error")

func Error() {
	//fmt.Printf("error: %v", errNotFound)
	if f, err := Sqrt(-1); err != nil {
		fmt.Printf("Error :%f / %s\n", f, err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return f, errors.New("math - square root of negative number")
	}
	return math.Sqrt(f), nil
}
