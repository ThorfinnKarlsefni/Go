package exercise

import (
	"fmt"
	"strconv"
)

type Celsius float64

func Cel() {
	var c Celsius = 18.36
	fmt.Println(c)
}

func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 32) + "°C"
}
