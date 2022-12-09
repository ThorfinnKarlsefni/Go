package exercise

import "fmt"

type Day int

const (
	MO Day = iota
	TU
	WE
	TH
	FR
	SA
	SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thurday", "Firday", "Saturday", "Sunday"}

func (day Day) String() string {
	return dayName[day]
}

func Days() {
	// var th Day = 6
	fmt.Println(6)
	fmt.Println(MO)
}
