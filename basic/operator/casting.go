package operator

import (
	"fmt"
	"math"
)

func Init() {
	var n int16 = 34
	var m int32

	m = int32(n)

	fmt.Printf("32 bit int is : %d\n", m)
	fmt.Printf("16 bit int is : %T\n", n)

}

//int converts int8

func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}

	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

// float64 converts int

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		//x lies int integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}

	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
