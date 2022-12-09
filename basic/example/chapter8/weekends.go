package chapter8

import "fmt"

func Weekend() {
	weekends := map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
	fmt.Printf("value is %s\n", weekends[2])
}
